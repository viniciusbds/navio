package container

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/docker/docker/pkg/reexec"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/logger"
	"github.com/viniciusbds/navio/utilities"
)

var l = logger.New(time.Kitchen, true)

func init() {
	reexec.Register("child", child)
	if reexec.Init() {
		os.Exit(0)
	}
}

// CreateContainer creates a container based on a baseImg, containerName and command with params
func CreateContainer(args []string) {
	baseImage, containerID, containerName, command, params := args[0], args[1], args[2], args[3], args[4:]
	prepareImage(baseImage, containerName)
	run(baseImage, containerID, containerName, command, params)
}

func prepareImage(baseImg, containerName string) {
	if !images.Exists(baseImg) {
		utilities.Must(images.Pull(baseImg))
	}
	if !images.RootfsExists(containerName) {
		images.PrepareRootfs(baseImg, containerName)
	}
	if baseImg == "ubuntu" {
		images.ConfigureNetworkForUbuntu(containerName)
	}
}

func run(baseImage string, containerID string, containerName string, command string, params []string) {
	container := NewContainer(containerID, containerName, baseImage, "-", containerName)
	InsertContainer(container)
	cmd := reexec.Command(append([]string{"child", baseImage, containerName, command}, params...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utilities.Must(cmd.Run())
}

func child() {
	_, containerName, command, params := os.Args[1], os.Args[2], os.Args[3], os.Args[4:]

	utilities.Must(syscall.Sethostname([]byte(containerName)))
	configureCgroups()
	pivotRoot(filepath.Join(utilities.RootfsPath, containerName))
	mountProc()

	cmd := exec.Command(command, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utilities.Must(cmd.Run())

	unmountProc()
}

// pivotRoot change the root file system of the process/container
// It moves the root file system of the current process to the
// directory putOld and makes newRoot the new root file system
// see more in https://linux.die.net/man/8/pivot_root
func pivotRoot(imagePath string) {
	newRoot := imagePath
	putOld := imagePath + "/put_old"
	utilities.Must(syscall.Mount(newRoot, newRoot, "", syscall.MS_BIND, ""))
	utilities.Must(os.MkdirAll(putOld, 0700))
	utilities.Must(syscall.PivotRoot(newRoot, putOld))
	utilities.Must(os.Chdir("/"))
	putOld = "/put_old"
	utilities.Must(syscall.Unmount(putOld, syscall.MNT_DETACH))
	utilities.Must(os.RemoveAll(putOld))
}

func mountProc() {
	if _, err := os.Stat("/proc"); os.IsNotExist(err) {
		os.Mkdir("/proc", 0700)
	}
	// source, target, fstype, flags, data
	err := syscall.Mount("proc", "/proc", "proc", 0, "")
	if err != nil {
		l.Log("ERROR", "The path /proc wasn't mounted: "+err.Error())
		os.Exit(1)
	}
}

func unmountProc() {
	// target, flags
	err := syscall.Unmount("/proc", 0)
	if err != nil {
		l.Log("ERROR", "The path /proc wasn't unmounted: "+err.Error())
		os.Exit(1)
	}
}

func configureCgroups() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "vini"), 0755)
	//fmt.Println(filepath.Join(pids, "vini"))
	utilities.Must(ioutil.WriteFile(filepath.Join(pids, "vini/pids.max"), []byte("24"), 0700))
	// Removes the new cgroup in place after the container exits
	utilities.Must(ioutil.WriteFile(filepath.Join(pids, "vini/notify_on_release"), []byte("1"), 0700))
	utilities.Must(ioutil.WriteFile(filepath.Join(pids, "vini/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

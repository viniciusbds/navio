package container

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/viniciusbds/navio/constants"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/pkg/logger"
	"github.com/viniciusbds/navio/pkg/reexec"
	"github.com/viniciusbds/navio/pkg/util"
)

var l = logger.New(time.Kitchen, true)

func init() {
	reexec.Register("child", child)
	if reexec.Init() {
		os.Exit(0)
	}
}

// CreateContainer creates a container based on a baseImg, containerName and command with params
func CreateContainer(containerID, containerName, baseImage, command string, params []string, prepare chan bool) error {
	prepareImage(baseImage, containerID)
	saveContainer(baseImage, containerID, containerName, command, params)
	prepare <- true
	return run(containerID, containerName, command, params)
}

func prepareImage(baseImg, containerID string) {
	if !images.Exists(baseImg) {
		util.Must(images.Pull(baseImg))
	}
	if !rootFSExists(containerID) {
		images.PrepareRootFS(baseImg, containerID)
	}
	if baseImg == "ubuntu" {
		images.ConfigureNetworkForUbuntu(containerID)
	}
}

func saveContainer(baseImage string, containerID string, containerName string, command string, params []string) {
	container := &Container{
		ID:      containerID,
		Name:    containerName,
		Image:   baseImage,
		Status:  "-",
		Root:    filepath.Join(constants.RootFSPath, containerName),
		Command: command,
		Params:  params,
	}
	Insert(container)
}

func run(containerID, containerName, command string, params []string) error {
	cmd := reexec.Command(append([]string{"child", containerID, containerName, command}, params...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	updateStatus(containerID, "Running")
	err := cmd.Run()
	updateStatus(containerID, "Stopped")
	return err
}

func child() {
	containerID, containerName, command, params := os.Args[1], os.Args[2], os.Args[3], os.Args[4:]
	util.Must(syscall.Sethostname([]byte(containerName)))
	configureCgroups()
	pivotRoot(filepath.Join(constants.RootFSPath, containerID))
	mountProc()
	cmd := exec.Command(command, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	util.Must(cmd.Run())
	unmountProc()
}

// pivotRoot change the root file system of the process/container
// It moves the root file system of the current process to the
// directory putOld and makes newRoot the new root file system
// see more in https://linux.die.net/man/8/pivot_root
func pivotRoot(imagePath string) {
	newRoot := imagePath
	putOld := imagePath + "/put_old"
	util.Must(syscall.Mount(newRoot, newRoot, "", syscall.MS_BIND, ""))
	util.Must(os.MkdirAll(putOld, 0700))
	util.Must(syscall.PivotRoot(newRoot, putOld))
	util.Must(os.Chdir("/"))
	putOld = "/put_old"
	util.Must(syscall.Unmount(putOld, syscall.MNT_DETACH))
	util.Must(os.RemoveAll(putOld))
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
	util.Must(ioutil.WriteFile(filepath.Join(pids, "vini/pids.max"), []byte("24"), 0700))
	// Removes the new cgroup in place after the container exits
	util.Must(ioutil.WriteFile(filepath.Join(pids, "vini/notify_on_release"), []byte("1"), 0700))
	util.Must(ioutil.WriteFile(filepath.Join(pids, "vini/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

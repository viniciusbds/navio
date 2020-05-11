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

// CreateContainer creates a container. Receive as argument: ["run", <image-name>, <command>, <params> ]
// [TODO]: Better document this function
func CreateContainer(args []string) {
	image, command, containerName, params := args[0], args[1], args[2], args[3:]
	prepareImage(image, containerName)
	run(image, command, containerName, params)
}

func prepareImage(image, containerName string) {
	if !images.TarImageExists(image) {
		images.Pull(image)
	}
	if !images.ImageIsReady(containerName) {
		images.Prepare(image, containerName)
	}
	if image == "ubuntu" {
		images.ConfigureNetworkForUbuntu(containerName)
	}
}

func run(image string, command string, containerName string, params []string) {
	cmd := reexec.Command(append([]string{"child", image, command, containerName}, params...)...)

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
	_, command, containerName, params := os.Args[1], os.Args[2], os.Args[3], os.Args[4:]

	utilities.Must(syscall.Sethostname([]byte(containerName)))
	configureCgroups()
	pivotRoot(filepath.Join(utilities.ImagesPath, containerName))
	mountProc()

	cmd := exec.Command(command, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utilities.Must(cmd.Run())

	unmountProc()
}

func pivotRoot(imagePath string) {
	oldrootfs := imagePath + "/oldrootfs"
	utilities.Must(syscall.Mount(imagePath, imagePath, "", syscall.MS_BIND, ""))
	utilities.Must(os.MkdirAll(oldrootfs, 0700))
	utilities.Must(syscall.PivotRoot(imagePath, oldrootfs))
	utilities.Must(os.Chdir("/"))
	oldrootfs = "/oldrootfs"
	utilities.Must(syscall.Unmount(oldrootfs, syscall.MNT_DETACH))
	utilities.Must(os.RemoveAll(oldrootfs))
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

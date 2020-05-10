package container

import (
	"fmt"
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
	util "github.com/viniciusbds/navio/utilities"
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

	if args[0] != "run" {
		l.Log("ERROR", "Bad command")
		os.Exit(1)
	}

	image, command, params := args[1], args[2], args[3:]

	if !images.AlreadyExists(image) {
		l.Log("WARNING", fmt.Sprintf("Image %s is not available, pull it ...", image))
		images.Pull(image)
	}

	run(image, command, params)
}

func run(image string, command string, params []string) {
	cmd := reexec.Command(append([]string{"child", image, command}, params...)...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	util.Must(cmd.Run())
}

func child() {
	l.Log("INFO", "Namespace setup code goes here <<\n")
	childRun(os.Args[1], os.Args[2], os.Args[3:])
}

func childRun(image string, command string, params []string) {
	util.Must(syscall.Sethostname([]byte("container")))

	configureCgroups()
	pivotRoot(utilities.ImagesRootDir + "/images/" + image)
	mountProc()

	cmd := exec.Command(command, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	util.Must(cmd.Run())

	unmountProc()
}

func pivotRoot(imagePath string) {
	oldrootfs := imagePath + "/oldrootfs"
	util.Must(syscall.Mount(imagePath, imagePath, "", syscall.MS_BIND, ""))
	util.Must(os.MkdirAll(oldrootfs, 0700))
	util.Must(syscall.PivotRoot(imagePath, oldrootfs))
	util.Must(os.Chdir("/"))
	oldrootfs = "/oldrootfs"
	util.Must(syscall.Unmount(oldrootfs, syscall.MNT_DETACH))
	util.Must(os.RemoveAll(oldrootfs))
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

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
	"github.com/viniciusbds/navio/src/logger"
	"github.com/viniciusbds/navio/src/util"
)

var l = logger.New(time.Kitchen, true)

func init() {
	reexec.Register("child", child)
	if reexec.Init() {
		os.Exit(0)
	}
}

func child() {
	l.Log("INFO", "Namespace setup code goes here <<\n")
	childRun(os.Args[1], os.Args[2], os.Args[3:])
}

func childRun(image string, command string, params []string) {

	configureCgroups()

	util.Must(syscall.Sethostname([]byte("container")))

	imagePath := "./images/" + image

	util.Must(syscall.Mount(imagePath, imagePath, "", syscall.MS_BIND, ""))
	util.Must(os.MkdirAll(imagePath+"/oldrootfs", 0700))
	util.Must(syscall.PivotRoot(imagePath, imagePath+"/oldrootfs"))
	util.Must(os.Chdir("/"))

	util.Must(mountProc(""))

	cmd := exec.Command(command, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	util.Must(cmd.Run())
	util.Must(unmountProc(""))
}

// CreateContainer creates a container. Receive as argument: ["run", <image-name>, <command>, <params> ]
// [TODO]: Better document this function
func CreateContainer(args []string) {

	if args[0] != "run" {
		l.Log("ERROR", "Bad command")
		os.Exit(1)
	}

	image, command, params := args[1], args[2], args[3:]

	if !images.CheckIfImageExists(image) {
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

func mountProc(newroot string) error {
	source := "proc"
	target := filepath.Join(newroot, "/proc")
	fstype := "proc"
	flags := 0
	data := ""

	_ = os.Mkdir(target, 0700)

	return syscall.Mount(source, target, fstype, uintptr(flags), data)
}

func unmountProc(newroot string) error {
	target := filepath.Join(newroot, "/proc")
	flags := 0
	return syscall.Unmount(target, flags)
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

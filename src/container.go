package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// docker         run image <cmd> <params>
// go run main.go run image <cmd> <params>
// navio container image <cmd> <params>

// CreateContainer creates a container
// Params: args
//          args[0] (run, child)
//          args[1] command
//          args[2:]... params
func CreateContainer(args []string) {
	//fmt.Println(args)
	switch args[0] {
	case "run":
		run(args[1], args[2:])
	case "child":
		child(args[1], args[2:])
	default:
		panic("bad command")
	}
}

func run(command string, params []string) {
	cmd := exec.Command("./navio", append([]string{"container", "child", command}, params...)...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	must(cmd.Run())
}

func child(command string, params []string) {
	//fmt.Printf("Running child %v as %d\n", os.Args[2:], os.Getpid())

	configureCgroups()

	syscall.Sethostname([]byte("container"))

	syscall.Chroot("../images/ubuntu")
	os.Chdir("/")

	_ = os.Mkdir("proc", 0700)

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(command, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
	must(syscall.Unmount("proc", 0))

}

func configureCgroups() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "vini"), 0755)
	//fmt.Println(filepath.Join(pids, "vini"))
	must(ioutil.WriteFile(filepath.Join(pids, "vini/pids.max"), []byte("24"), 0700))
	// Removes the new cgroup in place after the container exits
	must(ioutil.WriteFile(filepath.Join(pids, "vini/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "vini/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

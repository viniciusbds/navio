package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// docker         run image <cmd> <params>
// go run main.go run image <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run() {

	fmt.Printf("Running pai %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running child %v as %d\n", os.Args[2:], os.Getpid())

	cg()

	syscall.Sethostname([]byte("container"))

	syscall.Chroot("../images/ubuntu")
	os.Chdir("/")

	_ = os.Mkdir("proc", 0700)

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
	must(syscall.Unmount("proc", 0))

}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "vini"), 0755)
	fmt.Println(filepath.Join(pids, "vini"))
	must(ioutil.WriteFile(filepath.Join(pids, "vini/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	must(ioutil.WriteFile(filepath.Join(pids, "vini/notify_on_release"), []byte("1"), 0700))
	//must(ioutil.WriteFile(filepath.Join(pids, "vini/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

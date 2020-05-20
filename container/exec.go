package container

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
	"github.com/viniciusbds/navio/utilities"
)

func init() {
	reexec.Register("childExec", childExec)
	if reexec.Init() {
		os.Exit(0)
	}
}

// Exec executes a container based on a baseImg, containerName and command with params
func Exec(args []string) {
	baseImage, containerName, command, params := args[0], args[1], args[2], args[3:]
	prepareImage(baseImage, containerName)
	exe(containerName, command, params)
}

func exe(containerName string, command string, params []string) {
	cmd := reexec.Command(append([]string{"childExec", containerName, command}, params...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utilities.Must(cmd.Run())
}

func childExec() {
	containerName, command, params := os.Args[1], os.Args[2], os.Args[3:]
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

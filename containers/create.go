package containers

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
func CreateContainer(containerID, containerName, baseImage, command string, params []string, prepare chan bool, cgroup *CGroup) error {
	prepareImage(baseImage, containerID)
	saveContainer(baseImage, containerID, containerName, command, params, cgroup)
	prepare <- true
	return run(containerID, containerName, command, params)
}

func prepareImage(baseImg, containerID string) {
	if !images.IsAvailable(baseImg) {
		util.Must(images.Pull(baseImg))
	}
	if !rootFSExists(containerID) {
		images.PrepareRootFS(baseImg, containerID)
	}
	if baseImg == "ubuntu" {
		images.ConfigureNetworkForUbuntu(containerID)
	}
	if baseImg == "alpine" {
		// [TODO] images.ConfigureNetworkForAlpine(containerID)
	}
	if baseImg == "busybox" {
		// [TODO] images.ConfigureNetworkForBusybox(containerID)
	}
}

func saveContainer(baseImage string, containerID string, containerName string, command string, params []string, cgroup *CGroup) {
	container := NewContainer(containerID, containerName, baseImage, "-", filepath.Join(constants.RootFSPath, containerID), command, params, cgroup)
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
	container := getContainer(containerID)

	util.Must(syscall.Sethostname([]byte(containerName)))

	limitProcessCreation(container)
	limitCpus(container)
	limitCpushares(container)
	limitMemory(container)

	pivotRoot(container.RootFS)
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

func limitProcessCreation(container *Container) {
	cgroup := "/sys/fs/cgroup/"
	containerPids := filepath.Join(cgroup, "pids", container.ID)
	os.Mkdir(containerPids, 0755)

	maxpids := container.GetMaxpids()
	if maxpids == "" {
		maxpids = constants.DefaultMaxProcessCreation
	}

	util.Must(ioutil.WriteFile(filepath.Join(containerPids, "pids.max"), []byte(maxpids), 0700))
	// Removes the new cgroup in place after the container exits
	util.Must(ioutil.WriteFile(filepath.Join(containerPids, "notify_on_release"), []byte("1"), 0700))
	// Attach the process on the cgroup.
	util.Must(ioutil.WriteFile(filepath.Join(containerPids, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func limitCpus(container *Container) {
	cgroup := "/sys/fs/cgroup/"
	containerCPUSet := filepath.Join(cgroup, "cpuset", container.ID)
	os.Mkdir(containerCPUSet, 0755)

	cpus := container.GetCPUS()
	if cpus == "" {
		cpus = constants.DefaultCPUS
	}
	// 0 means use all memory for this cgroup
	util.Must(ioutil.WriteFile(filepath.Join(containerCPUSet, "cpuset.mems"), []byte("0"), 0700))
	util.Must(ioutil.WriteFile(filepath.Join(containerCPUSet, "cpuset.cpus"), []byte(cpus), 0700))
	// Removes the new cgroup in place after the container exits
	util.Must(ioutil.WriteFile(filepath.Join(containerCPUSet, "notify_on_release"), []byte("1"), 0700))
	// Attach the process on the cgroup.
	util.Must(ioutil.WriteFile(filepath.Join(containerCPUSet, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func limitCpushares(container *Container) {
	cgroup := "/sys/fs/cgroup/"
	containerCPUShares := filepath.Join(cgroup, "cpu", container.ID)
	os.Mkdir(containerCPUShares, 0755)

	cpushares := container.GetCPUshares()
	if cpushares == "" {
		cpushares = constants.DefaultCPUshares
	}

	util.Must(ioutil.WriteFile(filepath.Join(containerCPUShares, "cpu.shares"), []byte(cpushares), 0700))
	// Removes the new cgroup in place after the container exits
	util.Must(ioutil.WriteFile(filepath.Join(containerCPUShares, "notify_on_release"), []byte("1"), 0700))
	// Attach the process on the cgroup.
	util.Must(ioutil.WriteFile(filepath.Join(containerCPUShares, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func limitMemory(container *Container) {
	cgroup := "/sys/fs/cgroup/"
	containerMemory := filepath.Join(cgroup, "memory", container.ID)
	os.Mkdir(containerMemory, 0755)

	maxmemmory := container.GetMemory()
	if maxmemmory == "" {
		maxmemmory = constants.DefaultMemlimit
	}

	util.Must(ioutil.WriteFile(filepath.Join(containerMemory, "memory.limit_in_bytes"), []byte(maxmemmory), 0700))
	// Removes the new cgroup in place after the container exits
	util.Must(ioutil.WriteFile(filepath.Join(containerMemory, "notify_on_release"), []byte("1"), 0700))
	// Attach the process on the cgroup.
	util.Must(ioutil.WriteFile(filepath.Join(containerMemory, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

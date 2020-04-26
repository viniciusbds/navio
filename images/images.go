package images

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/viniciusbds/navio/src/logger"
	"github.com/viniciusbds/navio/src/util"
)

var (
	l = logger.New(time.Kitchen, true)

	// URL's
	alpineURL      = "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz"
	busyboxURL     = "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar"
	ubuntuFocalURL = "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz"

	// FileName's
	alpine      = "alpine-minirootfs-3.11.6-x86_64.tar.gz"
	busybox     = "busybox.tar"
	ubuntuFocal = "ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz"
)

// Pull ...
func Pull(imageName string) {

	if checkIfImageExists(imageName) {
		msg := fmt.Sprintf("The image %s already was downloaded", imageName)
		l.Log("WARNING", msg)
		return
	}

	var url, file string
	switch imageName {
	case "busybox":
		url = busyboxURL
		file = busybox
	case "ubuntu":
		url = ubuntuFocalURL
		file = ubuntuFocal
	}

	imagePath := fmt.Sprintf("images/%s", imageName)

	l.Log("INFO", fmt.Sprintf("Downloading %s  from %s ...", file, url))
	wgetCmd := exec.Command("wget", url)
	mkdirCmd := exec.Command("mkdir", "-p", imagePath)
	tarCmd := exec.Command("tar", "-C", imagePath, "-xf", file)
	rmFileCmd := exec.Command("rm", file)

	util.Must(wgetCmd.Run())
	util.Must(mkdirCmd.Run())
	util.Must(tarCmd.Run())
	util.Must(rmFileCmd.Run())

}

func checkIfImageExists(imageName string) bool {
	if _, err := os.Stat(fmt.Sprintf("./images/%s", imageName)); !os.IsNotExist(err) {
		return true
	}
	return false
}

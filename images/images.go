package images

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/src/logger"
	"github.com/viniciusbds/navio/src/util"
)

var (
	l       = logger.New(time.Kitchen, true)
	magenta = ansi.ColorFunc("magenta+")

	// URL's
	alpineURL  = "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz"
	busyboxURL = "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar"
	ubuntuURL  = "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz"
)

// Pull ...
// [TODO]: Document this function
func Pull(imageName string) {

	if CheckIfImageExists(imageName) {
		msg := fmt.Sprintf("The image %s already was downloaded", imageName)
		l.Log("WARNING", msg)
		return
	}

	var imageURL string
	switch imageName {
	case "alpine":
		imageURL = alpineURL
	case "busybox":
		imageURL = busyboxURL
	case "ubuntu":
		imageURL = ubuntuURL
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", imageName, imageURL))

	imagePath := fmt.Sprintf("./images/%s", imageName)

	if err := util.Wget(imageURL, imageName+".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled", imageName))
	}

	if err := os.Mkdir(imagePath, 0777); err != nil {
		l.Log("ERROR", fmt.Sprintf("The directory %s was not created", imagePath))
	}

	if err := util.Tar(imagePath, imageName+".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The file %s was not extracted", imageName+".tar"))
	}

	if err := os.Remove(imageName + ".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The file %s was not removed", imageName))
	}

	l.Log("INFO", "Pulled successfully :)\n")
}

// CheckIfImageExists ...
// [TODO]: Document this function
func CheckIfImageExists(imageName string) bool {
	if _, err := os.Stat(fmt.Sprintf("./images/%s", imageName)); !os.IsNotExist(err) {
		return true
	}
	return false
}

// ShowDownloadedImages ...
// [TODO]: Document this function
func ShowDownloadedImages() {
	dirname := "./images"

	f, err := os.Open(dirname)
	if err != nil {
		l.Log("ERROR", err.Error())
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		l.Log("ERROR", err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(magenta(file.Name()))
		}
	}
}

// RemoveDownloadedImages ...
// [TODO]: Document this function
func RemoveDownloadedImages(image string) {
	if CheckIfImageExists(image) {
		err := os.RemoveAll("./images/" + image)
		if err != nil {
			l.Log("ERROR", err.Error())
		} else {
			l.Log("INFO", fmt.Sprintf("The image %s was removed sucessfully!", image))
		}
	} else {
		l.Log("WARNING", fmt.Sprintf("The image %s doesn't exist.", image))
	}
}

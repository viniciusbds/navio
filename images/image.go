package images

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"errors"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/assert"
	"github.com/viniciusbds/navio/logger"
	"github.com/viniciusbds/navio/utilities"
)

var (
	l       = logger.New(time.Kitchen, true)
	magenta = ansi.ColorFunc("magenta+")
)

// Pull Downloads the .tar file from the official site
func Pull(imageName string) error {
	image := getImage(imageName)

	if image == nil {
		msg := "Invalid image: " + image.name
		l.Log("WARNING", msg)
		return errors.New(msg)
	}

	if TarImageExists(image.name) {
		l.Log("WARNING", "The image "+image.name+" already was downloaded")
		return nil
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", image.name, image.url))

	dir, _ := os.Getwd()
	if tarsPathExists := utilities.FileExists(utilities.TarsPath); !tarsPathExists {
		utilities.Must(os.MkdirAll(utilities.TarsPath, 0777))
	}
	utilities.Must(os.Chdir(utilities.TarsPath))
	err := utilities.Wget(image.url, image.name+".tar")
	if err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled\n %s", image.name, err.Error()))
		return err
	}
	utilities.Must(os.Chdir(dir))

	l.Log("INFO", "Pulled successfully :)\n")
	return nil
}

// Prepare receive as argument the imageName and the containerName, create a directory with the
// containerName and untar the respective image to this directory
func Prepare(imageName, containerName string) error {
	imagePath := filepath.Join(utilities.ImagesPath, containerName)
	tarFile := filepath.Join(utilities.TarsPath, imageName) + ".tar"
	if err := os.MkdirAll(imagePath, 0777); err != nil {
		l.Log("ERROR", fmt.Sprintf("The directory %s was not created \n%s", imagePath, err.Error()))
		return err
	}
	if err := utilities.Untar(imagePath, tarFile); err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not extracted. \n%s", imageName, err.Error()))
		return err
	}
	// fmt.Printf("before insert %v\n", availableImages)
	// availableImages = InsertANewImage(containerName, imageName, availableImages)
	// fmt.Printf("after insert %v\n", availableImages)

	return nil
}

// ConfigureNetworkForUbuntu Add the run/systemd/resolve/stub-resolv.conf file with the value "nameserver 8.8.8.8"
// see for more details: https://askubuntu.com/questions/91543/apt-get-update-fails-to-fetch-files-temporary-failure-resolving-error
func ConfigureNetworkForUbuntu(containerName string) {
	imagePath := filepath.Join(utilities.ImagesPath, containerName)
	resolveFile := filepath.Join(imagePath, "/run/systemd/resolve/stub-resolv.conf")
	if _, err := os.Stat(resolveFile); os.IsNotExist(err) {
		utilities.Must(os.MkdirAll(imagePath+"/run/systemd/resolve", 0777))
		//add a known DNS server to your system
		utilities.Must(ioutil.WriteFile(resolveFile, []byte("nameserver 8.8.8.8\n"), 0644))
	}
}

// ImageIsReady receive a containerName as argument and return TRUE if her rootfs is ready
// (i.e.: verify if there is a containerName dir on the ImagesPath directory)
func ImageIsReady(containerName string) bool {
	if _, err := os.Stat(filepath.Join(utilities.ImagesPath, containerName)); os.IsNotExist(err) {
		return false
	}
	return true
}

// TarImageExists receive a imageName as argument and return TRUE if the imageName.tar file exists
// on the default TarsPath directory (see it on utilities.contants)
func TarImageExists(imageName string) bool {
	if _, err := os.Stat(filepath.Join(utilities.TarsPath, imageName) + ".tar"); os.IsNotExist(err) {
		return false
	}
	return true
}

// ShowDownloadedImages ...
// [TODO]: Document this function
func ShowDownloadedImages() (string, error) {
	f, err := os.Open(utilities.ImagesPath)
	if err != nil {
		os.Mkdir(utilities.ImagesPath, 0777)
		return "", nil
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		l.Log("ERROR", err.Error())
		return "", err
	}

	var imageStr string
	result := ""
	for _, file := range files {
		if file.IsDir() {
			imageStr = file.Name() //getImage(file.Name()).ToStr()
			result += "\n" + magenta(imageStr)
		}
	}
	return result, nil
}

// DeleteImage ...
// [TODO]: Document this function
func DeleteImage(containerName string) error {
	if err := assert.ImageisNotEmpty(containerName); err != nil {
		return err
	}
	if ImageIsReady(containerName) {
		err := os.RemoveAll(filepath.Join(utilities.ImagesPath, containerName))
		if err != nil {
			l.Log("ERROR", err.Error())
			return err
		}
		l.Log("INFO", fmt.Sprintf("The image %s was removed sucessfully!", containerName))
	} else {
		l.Log("WARNING", fmt.Sprintf("The image %s doesn't exist.", containerName))
	}
	return nil
}

// Describe ...
func Describe(imageName string) string {
	image := getImage(imageName)
	if image == nil {
		l.Log("WARNING", fmt.Sprintf("Invalid image! Cannot describe %s", imageName))
		return ""
	}
	return getImage(imageName).ToStr()
}

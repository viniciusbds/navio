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
	if !utilities.IsaBaseImage(imageName) {
		msg := fmt.Sprintf("%s is not a base Image. Select one of the: %v", imageName, utilities.BaseImages)
		l.Log("WARNING", msg)
		return errors.New(msg)
	}

	image := getImage(imageName)

	if TarImageExists(image.name) {
		msg := "The image " + image.name + " already was downloaded"
		l.Log("WARNING", msg)
		return errors.New(msg)
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
	InsertImage(containerName, imageName)
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

// ShowBaseImages return a string with all official images
func ShowBaseImages() (string, error) {
	result := ""
	for _, img := range baseImages {
		result += "\n" + magenta(img.ToStr())
	}
	return result, nil
}

// Ps return a string with all availables container images that was created. see it like "containers"
func Ps() (string, error) {
	result := ""
	for _, img := range images {
		if !utilities.IsaBaseImage(img.name) {
			result += "\n" + magenta(img.ToStr())
		}
	}
	return result, nil
}

// DeleteImage receives a containerImage and remove it
func DeleteImage(containerName string) {
	if utilities.IsaBaseImage(containerName) {
		l.Log("WARNING", "Cannot remove the ["+containerName+"] base image")
		return
	}
	if err := assert.ImageisNotEmpty(containerName); err != nil {
		l.Log("WARNING", "Cannot remove a empty image: "+containerName)
		return
	}
	if ImageIsReady(containerName) {
		if err := os.RemoveAll(filepath.Join(utilities.ImagesPath, containerName)); err != nil {
			l.Log("ERROR", err.Error())
			return
		}
		l.Log("INFO", fmt.Sprintf("The image %s was removed sucessfully!", containerName))
	} else {
		l.Log("WARNING", fmt.Sprintf("The image %s doesn't exist.", containerName))
	}
	RemoveImage(containerName)
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

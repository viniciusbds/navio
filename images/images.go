package images

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"errors"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/logger"
	"github.com/viniciusbds/navio/utilities"
)

var (
	l       = logger.New(time.Kitchen, true)
	magenta = ansi.ColorFunc("magenta+")
)

// Pull Downloads the .tar file from the official site
func Pull(imageName string) error {
	if !utilities.IsOfficialImage(imageName) {
		return errors.New(imageName + " is not a official Image.")
	}

	image := getImage(imageName)

	if Exists(image.name) {
		return errors.New("The image " + image.name + " already was downloaded")
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", image.name, image.url))

	dir, _ := os.Getwd()
	if tarsPathExists := utilities.FileExists(utilities.ImagesPath); !tarsPathExists {
		utilities.Must(os.MkdirAll(utilities.ImagesPath, 0777))
	}
	utilities.Must(os.Chdir(utilities.ImagesPath))
	err := utilities.Wget(image.url, image.name+".tar")
	if err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled\n %s", image.name, err.Error()))
		return err
	}
	utilities.Must(os.Chdir(dir))

	l.Log("INFO", "Pulled successfully :)\n")
	return nil
}

// PrepareRootfs ...
func PrepareRootfs(baseImage, containerName string) error {
	rootfsPath := filepath.Join(utilities.RootfsPath, containerName)
	tarFile := filepath.Join(utilities.ImagesPath, baseImage) + ".tar"
	if err := os.MkdirAll(rootfsPath, 0777); err != nil {
		return err
	}
	if err := utilities.Untar(rootfsPath, tarFile); err != nil {
		return err
	}
	return nil
}

// ConfigureNetworkForUbuntu Add the run/systemd/resolve/stub-resolv.conf file with the value "nameserver 8.8.8.8"
// see for more details: https://askubuntu.com/questions/91543/apt-get-update-fails-to-fetch-files-temporary-failure-resolving-error
func ConfigureNetworkForUbuntu(containerName string) {
	rootfsPath := filepath.Join(utilities.RootfsPath, containerName)
	resolveFile := filepath.Join(rootfsPath, "/run/systemd/resolve/stub-resolv.conf")
	if _, err := os.Stat(resolveFile); os.IsNotExist(err) {
		utilities.Must(os.MkdirAll(rootfsPath+"/run/systemd/resolve", 0777))
		//add a known DNS server to your system
		utilities.Must(ioutil.WriteFile(resolveFile, []byte("nameserver 8.8.8.8\n"), 0644))
	}
}

// RootfsExists ...
func RootfsExists(containerName string) bool {
	if _, err := os.Stat(filepath.Join(utilities.RootfsPath, containerName)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Exists receive a imageName as argument and return TRUE if the imageName.tar file exists
// on the default TarsPath directory (see it on utilities.contants)
func Exists(image string) bool {
	if _, err := os.Stat(filepath.Join(utilities.ImagesPath, image) + ".tar"); os.IsNotExist(err) {
		return false
	}
	return true
}

// GetImages return a string with all base images
func GetImages() (result string, err error) {
	for _, img := range images {
		result += "\n" + magenta(img.ToStr())
	}
	return result, nil
}

// RemoveImage ...
func RemoveImage(image string) error {
	if utilities.IsOfficialImage(image) {
		return errors.New("Cannot remove a official image")
	}
	if utilities.IsEmpty(image) {
		return errors.New("Cannot remove a empty image")
	}
	if Exists(image) {
		if err := os.RemoveAll(filepath.Join(utilities.ImagesPath, image+".tar")); err != nil {
			return err
		}
	}
	deleteImage(image)
	return nil
}

func deleteImage(imgName string) {
	delete(images, imgName)
	removeImageDB(imgName)
}

// Describe ...
func Describe(imageName string) (string, error) {
	image := getImage(imageName)
	if image == nil {
		return "", errors.New("Invalid image! Cannot describe" + imageName)
	}
	return getImage(imageName).ToStr(), nil
}

// BuildANewBaseImg ...
func BuildANewBaseImg(name, baseImg string) error {
	newImgPath := filepath.Join(utilities.RootfsPath, name)
	tarFile := filepath.Join(utilities.ImagesPath, baseImg) + ".tar"
	if err := os.Mkdir(newImgPath, 0777); err != nil {
		return err
	}
	if err := utilities.Untar(newImgPath, tarFile); err != nil {
		return err
	}
	InsertImage(name, baseImg)
	return nil
}

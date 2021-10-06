package images

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/constants"
	"github.com/viniciusbds/navio/pkg/io"
	"github.com/viniciusbds/navio/pkg/logger"
	"github.com/viniciusbds/navio/pkg/util"
)

var (
	l       = logger.New(time.Kitchen, true)
	magenta = ansi.ColorFunc("magenta+")
	images  = make(map[string]*Image)
)

func init() {
	readImagesDB()
}

// Pull Downloads the .tar file from the official site
func Pull(imageName string) error {

	// This code ensures that the image (.tar file) is completely downloaded, if not, it is removed
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)

		// wait the signal
		<-sc

		// Remove the .tar file
		err := os.RemoveAll(filepath.Join(constants.ImagesPath, imageName+".tar"))
		if err != nil {
			fmt.Println(err)
		}

		os.Exit(0)
	}()

	if !constants.IsOfficialImage(imageName) {
		return errors.New(imageName + " is not a official Image.")
	}

	image := GetImage(imageName)

	if IsAvailable(image.Name) {
		return errors.New("The image " + image.Name + " already was downloaded")
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", image.Name, image.URL))

	dir, _ := os.Getwd()
	if tarsPathExists := io.FileExists(constants.ImagesPath); !tarsPathExists {
		util.Must(os.MkdirAll(constants.ImagesPath, 0777))
	}
	util.Must(os.Chdir(constants.ImagesPath))
	err := io.Wget(image.URL, image.Name+".tar")
	if err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled\n %s", image.Name, err.Error()))
		return err
	}
	util.Must(os.Chdir(dir))

	l.Log("INFO", "Pulled successfully :)\n")
	return nil
}

// PrepareRootFS ...
func PrepareRootFS(baseImage, containerID string) error {
	rootfsPath := filepath.Join(constants.RootFSPath, containerID)
	tarFile := filepath.Join(constants.ImagesPath, baseImage) + ".tar"
	if err := os.MkdirAll(rootfsPath, 0777); err != nil {
		return err
	}
	if err := io.Untar(rootfsPath, tarFile); err != nil {
		return err
	}
	return nil
}

// ConfigureNetworkForUbuntu Add the run/systemd/resolve/stub-resolv.conf file with the value "nameserver 8.8.8.8"
// see for more details: https://askubuntu.com/questions/91543/apt-get-update-fails-to-fetch-files-temporary-failure-resolving-error
func ConfigureNetworkForUbuntu(containerID string) {
	rootfsPath := filepath.Join(constants.RootFSPath, containerID)
	resolveFile := filepath.Join(rootfsPath, "/run/systemd/resolve/stub-resolv.conf")
	if _, err := os.Stat(resolveFile); os.IsNotExist(err) {
		util.Must(os.MkdirAll(rootfsPath+"/run/systemd/resolve", 0777))
		//add a known DNS server to your system
		util.Must(ioutil.WriteFile(resolveFile, []byte("nameserver 8.8.8.8\n"), 0644))
	}
}

// IsAvailable receive a imageName as argument and return TRUE if the imageName.tar file exists
// on the default TarsPath directory (see it on constants)
func IsAvailable(image string) bool {
	if _, err := os.Stat(filepath.Join(constants.ImagesPath, image) + ".tar"); os.IsNotExist(err) {
		return false
	}
	return true
}

// List return a string with all available images
func List() (result string) {
	for _, img := range images {
		result += "\n" + magenta(img.ToStr())
	}
	return
}

// Insert inserts a new image on the data structure and update the database
func Insert(name, baseImage string) error {
	baseImg := GetImage(baseImage)
	if baseImg == nil {
		return errors.New("ERROR: NIL Image ... ")
	}
	size, err := io.FileSize(baseImage)
	if err != nil {
		return err
	}
	baseImg.Size = strconv.FormatInt(size, 10)
	newImg := NewImage(name, baseImage, baseImg.Version, baseImg.Size, baseImg.URL)
	images[name] = newImg
	return insertImageDB(newImg)
}

// Remove a especific non official image
func Remove(name string) error {
	if constants.IsOfficialImage(name) {
		return errors.New("Cannot remove the " + name + " official image")
	}
	if util.IsEmpty(name) {
		return errors.New("Cannot remove a empty image")
	}
	if !IsAvailable(name) {
		return errors.New("Image " + name + " doesn't exist")
	}
	return removeImage(name)
}

// RemoveAll remove all non official images
func RemoveAll() error {
	for _, image := range images {
		if !constants.IsOfficialImage(image.Name) {
			err := removeImage(image.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetImage returns a image from the data structure
func GetImage(name string) *Image {
	return images[name]
}

// Untar extract the baseImage to create another one
func Untar(image, containerRootFS string, done chan bool) error {
	tarFile := filepath.Join(constants.ImagesPath, image) + ".tar"
	if err := os.Mkdir(containerRootFS, 0777); err != nil {
		return err
	}
	if err := io.Untar(containerRootFS, tarFile); err != nil {
		return err
	}
	done <- true
	return nil
}

func removeImage(name string) error {
	// First we remove the image.tar file
	err := os.RemoveAll(filepath.Join(constants.ImagesPath, name+".tar"))
	if err != nil {
		return err
	}
	// remove it from the data structure
	delete(images, name)
	// update it from the database
	removeImageDB(name)
	return nil
}

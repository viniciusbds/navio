package images

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/logger"
	"github.com/viniciusbds/navio/utilities"
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
	if !utilities.IsOfficialImage(imageName) {
		return errors.New(imageName + " is not a official Image.")
	}

	image := getImage(imageName)

	if Exists(image.Name) {
		return errors.New("The image " + image.Name + " already was downloaded")
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", image.Name, image.URL))

	dir, _ := os.Getwd()
	if tarsPathExists := utilities.FileExists(utilities.ImagesPath); !tarsPathExists {
		utilities.Must(os.MkdirAll(utilities.ImagesPath, 0777))
	}
	utilities.Must(os.Chdir(utilities.ImagesPath))
	err := utilities.Wget(image.URL, image.Name+".tar")
	if err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled\n %s", image.Name, err.Error()))
		return err
	}
	utilities.Must(os.Chdir(dir))

	l.Log("INFO", "Pulled successfully :)\n")
	return nil
}

// PrepareRootfs ...
func PrepareRootfs(baseImage, containerName string) error {
	rootfsPath := filepath.Join(utilities.RootFSPath, containerName)
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
	rootfsPath := filepath.Join(utilities.RootFSPath, containerName)
	resolveFile := filepath.Join(rootfsPath, "/run/systemd/resolve/stub-resolv.conf")
	if _, err := os.Stat(resolveFile); os.IsNotExist(err) {
		utilities.Must(os.MkdirAll(rootfsPath+"/run/systemd/resolve", 0777))
		//add a known DNS server to your system
		utilities.Must(ioutil.WriteFile(resolveFile, []byte("nameserver 8.8.8.8\n"), 0644))
	}
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
func GetImages() (result string) {
	for _, img := range images {
		result += "\n" + magenta(img.ToStr())
	}
	return
}

// InsertImage ...
func InsertImage(name, baseImage string) {
	baseImg := getImage(baseImage)
	newImg := NewImage(name, baseImage, baseImg.Version, baseImg.Size, baseImg.URL)
	images[name] = newImg
	insertImageDB(newImg)
}

// RemoveImage a especific non official image
func RemoveImage(name string) error {
	if utilities.IsOfficialImage(name) {
		return errors.New("Cannot remove the " + name + " official image")
	}
	if utilities.IsEmpty(name) {
		return errors.New("Cannot remove a empty image")
	}
	if !Exists(name) {
		return errors.New("Image " + name + " doesn't exist")
	}
	return removeImage(name)
}

// RemoveAllImages remove all non official images
func RemoveAllImages() {
	for _, image := range images {
		if !utilities.IsOfficialImage(image.Name) {
			removeImage(image.Name)
		}
	}
}

// Describe ...
func Describe(imageName string) (string, error) {
	image := getImage(imageName)
	if image == nil {
		return "", errors.New("Invalid image! Cannot describe" + imageName)
	}
	return image.ToStr(), nil
}

func getImage(name string) *Image {
	return images[name]
}

// UntarImg extract the baseImage to create another one
func UntarImg(name, baseImg string, done chan bool) error {
	newImgPath := filepath.Join(utilities.RootFSPath, name)
	tarFile := filepath.Join(utilities.ImagesPath, baseImg) + ".tar"
	if err := os.Mkdir(newImgPath, 0777); err != nil {
		return err
	}
	if err := utilities.Untar(newImgPath, tarFile); err != nil {
		return err
	}
	done <- true
	return nil
}

// IsValid receive a imageName and return true if is a valid image.
func IsValid(image string) bool {
	return getImage(image) != nil
}

func removeImage(name string) error {
	// remove it from the data structure
	delete(images, name)
	// update the database
	removeImageDB(name)
	// remove the .tar image file
	return os.RemoveAll(filepath.Join(utilities.ImagesPath, name+".tar"))
}

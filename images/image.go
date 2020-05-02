package images

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/assert"
	"github.com/viniciusbds/navio/logger"
	"github.com/viniciusbds/navio/utilities"
)

var (
	l       = logger.New(time.Kitchen, true)
	magenta = ansi.ColorFunc("magenta+")
)

// Pull ...
// [TODO]: Document this function
func Pull(imageName string) {
	image := getImage(imageName)

	if image == nil {
		msg := fmt.Sprintf("The image %s is not available", imageName)
		l.Log("WARNING", msg)
		return
	}

	if AlreadyExists(imageName) {
		msg := fmt.Sprintf("The image %s already was downloaded", imageName)
		l.Log("WARNING", msg)
		return
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", imageName, image.url))

	imagePath := "./images/" + imageName

	if err := utilities.Wget(image.url, imageName+".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled", imageName))
		l.Log("ERROR", fmt.Sprintf("%s", err.Error()))
	}

	if err := os.MkdirAll(imagePath, 0777); err != nil {
		l.Log("ERROR", fmt.Sprintf("The directory %s was not created", imagePath))
		l.Log("ERROR", fmt.Sprintf("%s", err.Error()))
	}

	if err := utilities.Tar(imagePath, imageName+".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The file %s was not extracted", imageName+".tar"))
		l.Log("ERROR", fmt.Sprintf("%s", err.Error()))
	}

	if err := os.Remove(imageName + ".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The file %s was not removed", imageName))
		l.Log("ERROR", fmt.Sprintf("%s", err.Error()))
	}

	l.Log("INFO", "Pulled successfully :)\n")
}

// AlreadyExists ...
// [TODO]: Document this function
func AlreadyExists(imageName string) bool {
	if _, err := os.Stat("./images/" + imageName); !os.IsNotExist(err) {
		return true
	}
	return false
}

// ShowDownloadedImages ...
// [TODO]: Document this function
func ShowDownloadedImages() string {
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

	var imageStr string
	result := "NAME\t\tVERSION\t\tSIZE"
	for _, file := range files {
		if file.IsDir() {
			imageStr = getImage(file.Name()).ToStr()
			result += "\n" + magenta(imageStr)
		}
	}
	return result
}

// RemoveDownloadedImage ...
// [TODO]: Document this function
func RemoveDownloadedImage(imageName string) error {
	if err := assert.ImageisNotEmpty(imageName); err != nil {
		return err
	}
	if AlreadyExists(imageName) {
		err := os.RemoveAll("./images/" + imageName)
		if err != nil {
			l.Log("ERROR", err.Error())
			return err
		}
		l.Log("INFO", fmt.Sprintf("The image %s was removed sucessfully!", imageName))
	} else {
		l.Log("WARNING", fmt.Sprintf("The image %s doesn't exist.", imageName))
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

package images

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
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

	if WasDownloaded(imageName) {
		msg := fmt.Sprintf("The image %s already was downloaded", imageName)
		l.Log("WARNING", msg)
		return
	}

	l.Log("INFO", fmt.Sprintf("Pulling %s  from %s ...", imageName, image.url))

	imagePath := "./images/" + imageName

	if err := utilities.Wget(image.url, imageName+".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The image %s was not Pulled", imageName))
	}

	if err := os.Mkdir(imagePath, 0777); err != nil {
		l.Log("ERROR", fmt.Sprintf("The directory %s was not created", imagePath))
	}

	if err := utilities.Tar(imagePath, imageName+".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The file %s was not extracted", imageName+".tar"))
	}

	if err := os.Remove(imageName + ".tar"); err != nil {
		l.Log("ERROR", fmt.Sprintf("The file %s was not removed", imageName))
	}

	l.Log("INFO", "Pulled successfully :)\n")
}

// WasDownloaded ...
// [TODO]: Document this function
func WasDownloaded(imageName string) bool {
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
func RemoveDownloadedImage(imageName string) {
	if WasDownloaded(imageName) {
		err := os.RemoveAll("./images/" + imageName)
		if err != nil {
			l.Log("ERROR", err.Error())
		} else {
			l.Log("INFO", fmt.Sprintf("The image %s was removed sucessfully!", imageName))
		}
	} else {
		l.Log("WARNING", fmt.Sprintf("The image %s doesn't exist.", imageName))
	}
}

// Describe ...
func Describe(imageName string) string {
	return "NAME\t\tVERSION\t\tSIZE\n" + getImage(imageName).ToStr()
}

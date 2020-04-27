package util

import (
	"os"
	"os/exec"
	"time"

	"github.com/viniciusbds/navio/src/logger"
)

var l = logger.New(time.Kitchen, true)

// Must ....
func Must(err error) {
	if err != nil {
		l.Log("ERROR", err.Error())
		os.Exit(1)
	}
}

// Contains ...
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Wget ...
func Wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Tar ...
func Tar(imagePath, imageName string) error {
	cmd := exec.Command("tar", "-C", imagePath, "-xf", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

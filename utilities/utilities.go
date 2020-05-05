package utilities

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/viniciusbds/navio/logger"
)

var l = logger.New(time.Kitchen, true)

// Must ....
// [TODO]: Document this function
func Must(err error) {
	if err != nil {
		l.Log("ERROR", err.Error())
		os.Exit(1)
	}
}

// Contains ...
// [TODO]: Document this function
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Wget ...
// [TODO]: Document this function
func Wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Tar ...
// [TODO]: Document this function
func Tar(imagePath, imageName string) error {
	cmd := exec.Command("tar", "-C", imagePath, "-xf", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// IsEmpty ...
func IsEmpty(imageName string) bool {
	if len(strings.TrimSpace(imageName)) == 0 {
		return true
	}
	return false
}

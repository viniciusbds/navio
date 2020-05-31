package utilities

import (
	"os"
	"os/exec"
	"strings"
	"sync"
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

// Untar ...
// [TODO]: Document this function
func Untar(directory, file string) error {
	cmd := exec.Command("tar", "-C", directory, "-xf", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Tar ...
func Tar(directory, file string) error {
	if err := os.Chdir(directory); err != nil {
		return err
	}
	cmd := exec.Command("tar", "cpjf", file, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Copy Copy a directory or a file from origen to a specific destiny
// (for ex: insidy the rootFS of a container)
func Copy(source, destiny string, wg *sync.WaitGroup) error {
	defer wg.Done()
	if !FileExists(destiny) {
		err := os.MkdirAll(destiny, 0777)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("cp", "-r", source, destiny)
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

// FileExists verifies if a directory or a file exists
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsOfficialImage ...
func IsOfficialImage(image string) bool {
	for _, i := range OfficialImages {
		if image == i {
			return true
		}
	}
	return false
}

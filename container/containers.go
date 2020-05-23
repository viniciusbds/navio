package container

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/utilities"
)

var (
	magenta    = ansi.ColorFunc("magenta+")
	containers = make(map[string]*Container)
)

func init() {
	readContainersDB()
}

// Ps return a string with all availables container images that was created. see it like "containers"
func Ps() (string, error) {
	result := ""
	for _, container := range containers {
		result += "\n" + magenta(container.ToStr())
	}
	return result, nil
}

// InsertContainer ...
func InsertContainer(container *Container) {
	containers[container.Name] = container
	insertContainersDB(container)
}

// RemoveContainer remove the rootfs directory of a container and remove it from the database.
func RemoveContainer(name string) error {
	if utilities.IsEmpty(name) {
		return errors.New("Empty container name")
	}
	if !exists(name) {
		return errors.New("Invalid container name")
	}
	// remove the rootFS
	if RootfsExists(name) {
		if err := os.RemoveAll(filepath.Join(utilities.RootFSPath, name)); err != nil {
			return err
		}
	}
	return deleteContainerRootfs(name)
}

// RootfsExists ...
func RootfsExists(containerName string) bool {
	if _, err := os.Stat(filepath.Join(utilities.RootFSPath, containerName)); os.IsNotExist(err) {
		return false
	}
	return true
}

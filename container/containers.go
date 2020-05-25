package container

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mgutz/ansi"
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

// InsertContainer insert a new container on the data structure and in the database
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
		return errors.New("Invalid container name: " + name)
	}
	// remove the rootFS
	if RootfsExists(name) {
		if err := os.RemoveAll(filepath.Join(utilities.RootFSPath, name)); err != nil {
			return err
		}
	}
	// remove the container from data structure
	delete(containers, name)
	// update the database
	return removeContainerDB(name)
}

func exists(name string) bool {
	return getContainer(name) != nil
}

func getContainer(name string) *Container {
	return containers[name]
}

// GetContainerID receive a container name and returns the respective ID
func GetContainerID(name string) string {
	result := ""
	c := getContainer(name)
	if c != nil {
		result = c.ID
	}
	return result
}

// GetContainerName receive a container ID and returns the respective name
func GetContainerName(ID string) string {
	result := ""
	for _, c := range containers {
		if c.ID == ID {
			result = c.Name
			return result
		}
	}
	return result
}

// RootfsExists ...
func RootfsExists(containerName string) bool {
	if _, err := os.Stat(filepath.Join(utilities.RootFSPath, containerName)); os.IsNotExist(err) {
		return false
	}
	return true
}

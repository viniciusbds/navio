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

// List return a string with all containers.
func List() (string, error) {
	result := ""
	for _, container := range containers {
		result += "\n" + magenta(container.ToStr())
	}
	return result, nil
}

// Insert inserts a new container on the data structure and update the database
func Insert(container *Container) {
	if container != nil {
		containers[container.Name] = container
		insertContainersDB(container)
	} else {
		l.Log("ERROR", "NIL container")
	}
}

// RemoveContainer remove a container by her ID
func RemoveContainer(ID string) error {
	if utilities.IsEmpty(ID) {
		return errors.New("Empty container ID")
	}
	if !Exists(ID) {
		return errors.New("Invalid container ID: " + ID)
	}
	if !RootfsExists(ID) {
		return errors.New("RootFS of container" + ID + " doesn't exist")
	}
	return removeContainer(ID)
}

// Exists receives a [containerName or containerID] and return true if the Container exists on the system
func Exists(arg string) bool {
	for _, container := range containers {
		if container.Name == arg || container.ID == arg {
			return true
		}
	}
	return false
}

// IsaID verifies if a string is the ID of some container
func IsaID(ID string) bool {
	for _, container := range containers {
		if container.ID == ID {
			return true
		}
	}
	return false
}

func getContainer(arg string) *Container {
	for _, container := range containers {
		if container.Name == arg || container.ID == arg {
			return containers[arg]
		}
	}
	return nil
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

func removeContainer(ID string) error {
	// remove it from the data structure
	delete(containers, ID)
	// update the database
	removeContainerDB(ID)
	// remove the rootFS
	return os.RemoveAll(filepath.Join(utilities.RootFSPath, ID))
}

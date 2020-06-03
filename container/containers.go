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
		containers[container.ID] = container
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
	if !RootFSExists(ID) {
		return errors.New("RootFS of container" + ID + " doesn't exist")
	}
	return removeContainer(ID)
}

// RemoveAll remove all containers
func RemoveAll() error {
	for _, container := range containers {
		err := removeContainer(container.ID)
		if err != nil {
			return err
		}
	}
	return nil
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
			return container
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

// RootFSExists receives a container ID and verifies if exists a RootFS
func RootFSExists(ID string) bool {
	_, err := os.Stat(filepath.Join(utilities.RootFSPath, ID))
	return !os.IsNotExist(err)
}

func removeContainer(ID string) error {
	// remove the rootFS
	err := os.RemoveAll(filepath.Join(utilities.RootFSPath, ID))
	if err != nil {
		return err
	}
	// remove it from the data structure
	delete(containers, ID)
	// update the database
	removeContainerDB(ID)
	return nil
}

// UpdateStatus update the Status of a Container
func updateStatus(ID, status string) error {
	if !IsaID(ID) {
		return errors.New("ERROR: Container not exists")
	}
	containers[ID].SetStatus(status)
	return updateContainerStatusDB(ID, status)
}

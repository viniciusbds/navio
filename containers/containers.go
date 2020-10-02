package containers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/mgutz/ansi"
	"github.com/viniciusbds/navio/constants"
	"github.com/viniciusbds/navio/pkg/util"
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

// Remove removes a container by her ID
func Remove(ID string) error {
	if util.IsEmpty(ID) {
		return errors.New("Empty container ID")
	}
	if !IsaID(ID) {
		return errors.New("Invalid container ID: " + ID)
	}
	if !rootFSExists(ID) {
		return errors.New("RootFS of container" + ID + " doesn't exist")
	}
	return removeContainer(ID)
}

// RemoveAll remove all containers
func RemoveAll(done chan bool) {
	for _, container := range containers {
		removeContainer(container.ID)
	}
	done <- true
}

// IsaID receives a containerID and return true if there's a container associated
func IsaID(ID string) bool {
	for _, container := range containers {
		if container.ID == ID {
			return true
		}
	}
	return false
}

func getContainer(ID string) *Container {
	return containers[ID]
}

// GetContainerID receive a container name and returns her respective ID
func GetContainerID(name string) string {
	result := ""
	for _, container := range containers {
		if container.Name == name {
			result = container.ID
			return result
		}
	}
	return result
}

// GetContainerName receive a container ID and returns her respective name
func GetContainerName(ID string) string {
	result := ""
	container := getContainer(ID)
	if container != nil {
		result = container.Name
	}
	return result
}

// UsedName receives a containerName and return true if the name already was used
func UsedName(name string) bool {
	ID := GetContainerID(name)
	return GetContainerName(ID) != ""
}

func rootFSExists(ID string) bool {
	_, err := os.Stat(filepath.Join(constants.RootFSPath, ID))
	return !os.IsNotExist(err)
}

func removeContainer(ID string) error {
	// First we try to remove the container rootFS
	err := os.RemoveAll(filepath.Join(constants.RootFSPath, ID))
	if err != nil {
		return err
	}
	// remove it from the data structure
	delete(containers, ID)
	// remove it from the database
	removeContainerDB(ID)
	return nil
}

func updateStatus(ID, status string) error {
	if !IsaID(ID) {
		return errors.New("ERROR: Container not exists")
	}
	containers[ID].SetStatus(status)
	return updateContainerStatusDB(ID, status)
}

func UpdateName(ID, name string) error {
	if !IsaID(ID) {
		return errors.New("ERROR: Container not exists")
	}
	containers[ID].SetName(name)
	return updateContainerNameDB(ID, name)
}

func numbeOfContainers() int {
	return len(containers)
}

// GenerateNewID gerenates a new random Container ID
func GenerateNewID() string {
	rand.Seed(time.Now().UnixNano())
	min := 11111111
	max := 99999999
	return fmt.Sprintf("%d", rand.Intn(max-min+1)+min)
}

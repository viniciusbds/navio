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

// RemoveContainerRootfs remove the rootfs directory of a container
func RemoveContainerRootfs(name string) error {
	if utilities.IsEmpty(name) {
		return errors.New("Empty container name")
	}
	if images.RootfsExists(name) {
		if err := os.RemoveAll(filepath.Join(utilities.RootfsPath, name)); err != nil {
			return err
		}
	}
	return deleteContainerRootfs(name)
}

func deleteContainerRootfs(name string) error {
	delete(containers, name)
	return removeContainerDB(name)
}

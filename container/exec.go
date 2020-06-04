package container

import (
	"errors"
)

// Exec runs an existing container.
func Exec(containerID, containerName, command string, params []string) error {
	if Exists(containerID) {
		return run(containerID, containerName, command, params)
	}
	return errors.New("The container " + containerName + " doesn't exist")
}

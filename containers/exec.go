package containers

import (
	"errors"
)

// Exec runs an existing container.
func Exec(containerID, command string, params []string) error {
	if IsaID(containerID) {
		containerName := GetContainerName(containerID)
		return run(containerID, containerName, command, params)
	}
	return errors.New("The container " + containerID + " doesn't exist")
}

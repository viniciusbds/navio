package container

import (
	"errors"
)

// Exec runs an existing container.
func Exec(args []string) error {
	containerID, containerName, command, params := args[0], args[1], args[2], args[3:]
	cont := getContainer(containerName)
	if cont == nil {
		return errors.New("The container " + containerName + " doesn't exist")
	}
	prepareImage(cont.Image, containerName)
	return run(containerID, containerName, command, params)
}

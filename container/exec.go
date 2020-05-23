package container

import (
	"errors"
)

// Exec runs an existing container.
func Exec(args []string) error {
	containerName, command, params := args[0], args[1], args[2:]
	cont := getContainer(containerName)
	if cont == nil {
		return errors.New("The container " + containerName + " doesn't exist")
	}
	prepareImage(cont.Image, containerName)
	return run(containerName, command, params)
}

package container

import (
	"testing"
)

func TestRemoveContainer(t *testing.T) {
	t.Run("Invalid containerName", func(t *testing.T) {
		name := "atata"
		err := RemoveContainer(name)
		if err != nil {
			result := err.Error()
			expected := "Invalid container name: " + name
			check(t, expected, result)
		}
	})
	t.Run("Empty name", func(t *testing.T) {
		err := RemoveContainer("    ")
		result := err.Error()
		expected := "Empty container name"
		check(t, expected, result)
	})
}

// [TODO: test] Ps return a string with all availables container images that was created. see it like "containers"

// [TODO: test] InsertContainer ...

// [TODO: test] RemoveContainer remove the rootfs directory of a container and remove it from the database.

// [TODO: test] GetContainerID receive a container name and returns the respective ID

// [TODO: test] GetContainerName receive a container ID and returns the respective name

// [TODO: test] RootfsExists ...

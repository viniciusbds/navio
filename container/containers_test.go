package container

import (
	"testing"
)

func TestRemoveContainerRootfs(t *testing.T) {

	t.Run("Invalid containerName", func(t *testing.T) {
		err := RemoveContainerRootfs("atata")
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("Empty name", func(t *testing.T) {
		err := RemoveContainerRootfs("    ")
		if err.Error() != "Empty container name" {
			t.Errorf("Expected a different msg")
		}
	})

}

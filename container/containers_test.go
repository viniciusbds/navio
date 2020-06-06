package container

import (
	"testing"
)

func TestRemoveContainer(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		id := "45252423432"
		err := RemoveContainer(id)
		result := err.Error()
		expected := "Invalid container ID: " + id
		check(t, expected, result)
	})
	t.Run("Empty ID", func(t *testing.T) {
		err := RemoveContainer("    ")
		result := err.Error()
		expected := "Empty container ID"
		check(t, expected, result)
	})
	t.Run("Valid ID", func(t *testing.T) {
		ID := "4351987343"

		go CreateContainer(ID, "gbn13am", "alpine", "echo", []string{"zizo"}, done)
		<-done

		if container := getContainer(ID); container == nil {
			t.Error("ERROR on Test RemoveContainer")
		}

		RemoveContainer(ID)

		if container := getContainer(ID); container != nil {
			t.Error("ERROR on Test RemoveContainer")
		}

		if IsaID(ID) {
			t.Error("ERROR on Test RemoveContainer")
		}
	})
}

// [TODO: test] InsertContainer ...

// [TODO: test] RemoveContainer remove the rootfs directory of a container and remove it from the database.

// [TODO: test] GetContainerID receive a container name and returns the respective ID

// [TODO: test] GetContainerName receive a container ID and returns the respective name

// [TODO: test] RootfsExists ...

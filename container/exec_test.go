package container

import (
	"testing"
)

var check = func(t *testing.T, expected string, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %s != Result %s", expected, result)
	}
}

var done chan bool

func init() {
	done = make(chan bool)
}

func TestExec(t *testing.T) {

	t.Run("Try EXEC a invalid container (i.e:  that doesn't exists)", func(t *testing.T) {
		containerName := "someIncommunnamiss"
		containerID := GetContainerID(containerName)
		err := Exec([]string{containerID, containerName, "echo", "Ola", "Menino", "Jesus", "de", "Atocha"})
		if err == nil {
			t.Errorf("Here we expected that err != nil, because the containerName doesn't exists!!!")
		}
		if err != nil {
			expected := "The container " + containerName + " doesn't exist"
			result := err.Error()
			check(t, expected, result)
		}
	})

	t.Run("EXEC a valid container (i.e:  that exists)", func(t *testing.T) {
		containerName := "galan-alpine"
		containerID := GetContainerID(containerName)

		// Cleaning before start
		if cont := getContainer(containerName); cont != nil {
			containerID := GetContainerID(containerName)
			RemoveContainer(containerID)
		}

		err := Exec([]string{containerID, containerName, "echo", "Ola", "Menino", "Jesus", "de", "Atocha"})

		if err != nil {
			expected := "The container " + containerName + " doesn't exist"
			result := err.Error()
			check(t, expected, result)
		}

		baseImg := "alpine"
		containerID = "98989898989"
		go CreateContainer([]string{baseImg, containerID, containerName, "echo", "creating container"}, done)
		<-done

		// Testing exec
		err = Exec([]string{containerID, containerName, "echo", "Ola", "Menino", "Jesus", "de", "Atocha"})
		if err != nil {
			t.Errorf(err.Error())
		}

		// Cleaning
		containerID = GetContainerID(containerName)
		RemoveContainer(containerID)
	})

}

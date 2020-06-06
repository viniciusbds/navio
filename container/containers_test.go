package container

import (
	"path/filepath"
	"testing"

	"github.com/viniciusbds/navio/constants"
)

func TestInsert(t *testing.T) {
	ID := GenerateNewID()
	cont := NewContainer(ID, "conteiir", "alpine", "Test", filepath.Join(constants.RootFSPath, "conteiir"), "echo", []string{"oi", "oi"})

	AssertContainerDontExists(ID, t)
	Insert(cont)
	AssertContainerExists(ID, t)
}

func TestRemoveContainer(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		id := GenerateNewID()
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
		ID := GenerateNewID()

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

func TestRemoveAll(t *testing.T) {
	go RemoveAll(done)
	<-done

	if numbeOfContainers() != 0 {
		t.Error("ERROR on Test RemoveAll: in begginer thes numbeofcontainers != 0")
	}

	id1, id2, id3 := GenerateNewID(), GenerateNewID(), GenerateNewID()
	go CreateContainer(id1, "name1", "alpine", "echo", []string{"echoed1"}, done)
	<-done
	go CreateContainer(id2, "name2", "alpine", "echo", []string{"echoed2"}, done)
	<-done
	go CreateContainer(id3, "name3", "alpine", "echo", []string{"echoed3"}, done)
	<-done
	if numbeOfContainers() != 3 {
		t.Error("ERROR on Test RemoveAll: we create 3 containers and the numbeofcontainers != 3")
	}

	for _, ID := range []string{id1, id2, id3} {
		AssertContainerExists(ID, t)
	}

	go RemoveAll(done)
	<-done

	for _, ID := range []string{id1, id2, id3} {
		AssertContainerDontExists(ID, t)
	}

	if numbeOfContainers() != 0 {
		t.Error("ERROR on Test RemoveAll: in end the numbeofcontainers != 0")
	}
}

// [TODO: test] IsaID receives a containerID and return true if there's a container associated
func TestIsaID(t *testing.T) {
	ID := GenerateNewID()

	result, expected := IsaID(ID), false
	checkbool(t, expected, result)

	go CreateContainer(ID, "gbn", "alpine", "echo", []string{"zizo"}, done)
	<-done

	result, expected = IsaID(ID), true
	checkbool(t, expected, result)
}

func TestGetContainerID(t *testing.T) {
	ID := GenerateNewID()
	go CreateContainer(ID, "paraybba", "alpine", "echo", []string{"campina grande é a city"}, done)
	<-done
	expected := ID
	result := GetContainerID("paraybba")
	check(t, expected, result)
}

func TestGetContainerName(t *testing.T) {
	ID := GenerateNewID()
	go CreateContainer(ID, "paraybba", "alpine", "echo", []string{"campina grande é a city"}, done)
	<-done
	expected := "paraybba"
	result := GetContainerName(ID)
	check(t, expected, result)
}

// [TODO: test] UsedName receives a containerName and return true if the name already was used
func TestUsedName(t *testing.T) {
	ID := GenerateNewID()
	go CreateContainer(ID, "paraybba", "alpine", "echo", []string{"campina grande é a city"}, done)
	<-done

	expected := true
	result := UsedName("paraybba")
	checkbool(t, expected, result)

	expected = false
	result = UsedName("Seu pereira e o coletivo   40ejul1")
	checkbool(t, expected, result)
}

var checkbool = func(t *testing.T, expected bool, result bool) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %v != Result %v", expected, result)
	}
}

func AssertContainerExists(ID string, t *testing.T) {
	if !IsaID(ID) {
		t.Errorf("ERROR: Container %s doesn't exists", ID)
	}
	if container := getContainer(ID); container == nil {
		t.Errorf("ERROR: Container %s IS NILL", ID)
	}
}

func AssertContainerDontExists(ID string, t *testing.T) {
	if IsaID(ID) {
		t.Errorf("ERROR: Container %s doesn't exists", ID)
	}
	if container := getContainer(ID); container != nil {
		t.Errorf("ERROR: Container %s IS NILL", ID)
	}
}

// [TODO: test] func getContainer(ID string)

// [TODO: test] func rootFSExists(ID string)

// [TODO: test] func removeContainer(ID string)

// [TODO: test] func updateStatus(ID, status string)

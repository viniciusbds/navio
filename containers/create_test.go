package containers

import (
	"testing"
)

func TestCreateContainer(t *testing.T) {

	id1, id2, id3 := GenerateNewID(), GenerateNewID(), GenerateNewID()
	for _, ID := range []string{id1, id2, id3} {
		AssertContainerDontExists(ID, t)
	}

	go CreateContainer(id1, "alpix", "alpine", "echo", []string{"echoedalpix"}, done, nil)
	<-done
	go CreateContainer(id2, "alpex", "alpine", "echo", []string{"echoedalpex"}, done, nil)
	<-done
	go CreateContainer(id3, "alpux", "alpine", "echo", []string{"echoedalpux"}, done, nil)
	<-done

	for _, ID := range []string{id1, id2, id3} {
		AssertContainerExists(ID, t)
	}
}

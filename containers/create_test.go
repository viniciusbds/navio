package containers

import (
	"testing"
)

func TestCreateContainer(t *testing.T) {

	id1, id2, id3 := GenerateNewID(), GenerateNewID(), GenerateNewID()
	for _, ID := range []string{id1, id2, id3} {
		AssertContainerDontExists(ID, t)
	}
	errs := make(chan error, 1)
	go func() {
		errs <- CreateContainer(id1, "alpix", "alpine", "echo", []string{"echoedalpix"}, done, nil)
	}()
	<-done
	if err := <-errs; err != nil {
		t.Errorf("ERROR on Test TestCreateContainer, %s", err)
	}

	go func() {
		errs <- CreateContainer(id2, "alpex", "alpine", "echo", []string{"echoedalpex"}, done, nil)
	}()
	<-done
	if err := <-errs; err != nil {
		t.Errorf("ERROR on Test TestCreateContainer, %s", err)
	}

	go func() {
		errs <- CreateContainer(id3, "alpux", "alpine", "echo", []string{"echoedalpux"}, done, nil)
	}()
	<-done
	if err := <-errs; err != nil {
		t.Errorf("ERROR on Test TestCreateContainer, %s", err)
	}

	for _, ID := range []string{id1, id2, id3} {
		AssertContainerExists(ID, t)
	}
}

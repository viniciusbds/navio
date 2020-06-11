package containers

import (
	"path/filepath"
	"testing"

	"github.com/viniciusbds/navio/constants"
)

func TestNewContainer(t *testing.T) {
	cont := NewContainer("98907342", "conteiir", "ubuntu", "OK", filepath.Join(constants.RootFSPath, "conteiir"), "echo", []string{"oi", "oi"}, nil)
	if cont.Name != "conteiir" || cont.Status != "OK" {
		t.Errorf("Coisas estranhas aconteceram")
	}
}

func TestIsRunning(t *testing.T) {
	cont := NewContainer("45216326", "conteiir", "ubuntu", "Stopped", filepath.Join(constants.RootFSPath, "conteiir"), "echo", []string{"oi", "oi"}, nil)
	result := cont.IsRunning()
	expected := false
	if expected != result {
		t.Errorf("Expected %v != Result %v", expected, result)
	}

	cont.SetStatus("Running")
	result = cont.IsRunning()
	expected = true
	if expected != result {
		t.Errorf("Expected %v != Result %v", expected, result)
	}
}

func TestSetStatus(t *testing.T) {
	cont := NewContainer("5425146", "conteiir", "ubuntu", "Running", filepath.Join(constants.RootFSPath, "conteiir"), "echo", []string{"oi", "oi"}, nil)
	cont.SetStatus("Stopped")
	result := cont.GetStatus()
	expected := "Stopped"
	check(t, expected, result)
}

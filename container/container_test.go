package container

import (
	"path/filepath"
	"testing"

	"github.com/viniciusbds/navio/constants"
)

func TestNewContainer(t *testing.T) {
	cont := NewContainer("98907342", "conteiir", "ubuntu", "OK", filepath.Join(constants.RootFSPath, "conteiir"), "echo", []string{"oi", "oi"})
	if cont.Name != "conteiir" || cont.Status != "OK" {
		t.Errorf("Coisas estranhas aconteceram")
	}
}

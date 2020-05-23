package images

import (
	"testing"
)

func TestNewImage(t *testing.T) {
	img := NewImage("ubunta", "ubuntu", "30.04", "50mb", "www.ubuntu.com")

	if img.Name != "ubunta" || img.Base != "ubuntu" {
		t.Error("Coisas estranhas aconteceram")
	}
}

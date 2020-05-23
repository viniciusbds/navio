package images

import (
	"testing"
)

var check = func(t *testing.T, expected string, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %s != Result %s", expected, result)
	}
}

func TestPull(t *testing.T) {
	//note: this tests don't cover all Pull function
	t.Run("Invalid Image", func(t *testing.T) {
		imageName := "ubuntuxxx"
		err := Pull(imageName)
		if err != nil {
			expected := imageName + " is not a official Image."
			result := err.Error()
			check(t, expected, result)
		}
	})
}

// [TODO: test] PrepareRootfs ...

// [TODO: test] ConfigureNetworkForUbuntu ...

// [TODO: test] Exists receive a imageName as argument and return TRUE if the imageName.tar file exists

func TestGetImages(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if _, err := GetImages(); err != nil {
			t.Errorf("ERROR: on GetImages(): %s", err.Error())
		}
	})
}

// [TODO: test] InsertImage ...

// [TODO: test] RemoveImage ...

// [TODO: test] Describe ...

// [TODO: test] BuildANewBaseImg ...

// [TODO: test] IsValid receive a imageName and return true if is a valid image.

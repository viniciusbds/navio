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

// [TODO: test] GetImages ...

// [TODO: test] InsertImage ...

func TestRemoveImage(t *testing.T) {

	// BASIC TESTS ------------------------------------------------------
	t.Run("Fail on remove a official Image", func(t *testing.T) {
		imageName := "ubuntu"
		err := RemoveImage(imageName)
		expected := "Cannot remove the " + imageName + " official image"
		result := err.Error()
		check(t, expected, result)
	})
	t.Run("Empty Image", func(t *testing.T) {
		imageName := "       "
		err := RemoveImage(imageName)
		expected := "Cannot remove a empty image"
		result := err.Error()
		check(t, expected, result)

	})
	t.Run("Image that doesn't exists", func(t *testing.T) {
		imageName := "ubuntuxxx"
		err := RemoveImage(imageName)
		expected := "Image " + imageName + " doesn't exist"
		result := err.Error()
		check(t, expected, result)

	})
	// END BASIC TESTS ------------------------------------------------------

	// ----------------------------------------------------------------- TODO
	t.Run("Remove a valid image", func(t *testing.T) {
		//imageName := "ubuntu-pyhtone"
		// CRIAR A IMAGEM COM O BUILD
		// VERIFICAR QUE ELA EXISTE
		// REMOVE-LA
	})
	// ----------------------------------------------------------------- TODO

}

// [TODO: test] Describe ...

// [TODO: test] BuildANewBaseImg ...

// [TODO: test] IsValid receive a imageName and return true if is a valid image.

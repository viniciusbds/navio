package images

import (
	"testing"

	"github.com/viniciusbds/navio/constants"
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

func TestRemoveImage(t *testing.T) {

	// BASIC TESTS ------------------------------------------------------
	t.Run("Fail on remove a official Image", func(t *testing.T) {
		imageName := "ubuntu"
		err := Remove(imageName)
		expected := "Cannot remove the " + imageName + " official image"
		result := err.Error()
		check(t, expected, result)
	})
	t.Run("Empty Image", func(t *testing.T) {
		imageName := "       "
		err := Remove(imageName)
		expected := "Cannot remove a empty image"
		result := err.Error()
		check(t, expected, result)

	})
	t.Run("Image that doesn't exists", func(t *testing.T) {
		imageName := "ubuntuxxx"
		err := Remove(imageName)
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

func AssertImageDontExists(name string, t *testing.T) {
	if IsAvailable(name) {
		t.Errorf("ERROR: here we expected that the image don't exists")
	}
	if image := GetImage(name); image != nil {
		t.Errorf("ERROR: here we expected a nil image. image: %s", image)
	}
}

func AssertImageExists(name string, t *testing.T) {
	if image := GetImage(name); image == nil {
		t.Error("ERROR: here we expected a non nil image.")
	}
}

func TestInsert(t *testing.T) {
	imageName := "novaaply"
	imageBase := "alpine"
	AssertImageDontExists(imageName, t)
	err := Insert(imageName, 2, imageBase)
	if err != nil {
		t.Errorf("Error: problem inserting image, %s", err)
	}
	AssertImageExists(imageName, t)

	// clear
	err = Remove(imageName)
	if err != nil {
		t.Errorf("Error: problem deleting image, %s", err)
	}

}

func TestIsAvailable(t *testing.T) {
	if !IsAvailable("alpine") {
		err := Pull("alpine")
		if err != nil {
			t.Errorf("Error: problem pulling image, %s", err)
		}

	}

	if !IsAvailable("busybox") {
		err := Pull("busybox")
		if err != nil {
			t.Errorf("Error: problem pulling image, %s", err)
		}

	}

	if !IsAvailable("ubuntu") {
		err := Pull("ubuntu")
		if err != nil {
			t.Errorf("Error: problem pulling image, %s", err)
		}

	}

	if !IsAvailable("alpine") || !IsAvailable("busybox") || !IsAvailable("ubuntu") {
		t.Error("ERROR: on Test IsAvailable")
	}
}

func TestRemoveAll(t *testing.T) {
	err := RemoveAll()
	if err != nil {
		t.Errorf("Error: problem deleting all images, %s", err)
	}

	if len(images) != len(constants.OfficialImages) {
		t.Errorf("ERROR on RemoveAll")
	}
}

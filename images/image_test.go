package images

import (
	"os"
	"testing"
)

var check = func(t *testing.T, expected string, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %s != Result %s", expected, result)
	}
}

var clear = func() {
	DeleteImage("alpine")
	DeleteImage("busybox")
	DeleteImage("ubuntu")
	os.Remove("images")
}

func TestPull(t *testing.T) {

	//note: this tests don't cover all Pull function

	t.Run("Invalid Image", func(t *testing.T) {
		imageName := "ubuntuxxx"
		err := Pull(imageName)
		if err != nil {
			expected := "The image " + imageName + " is not available"
			result := err.Error()
			check(t, expected, result)
		}
	})

	t.Run("A Image that already was downloaded", func(t *testing.T) {
		imageName := "alpine"
		Pull(imageName)

		err := Pull(imageName)
		if err != nil {
			t.Errorf("EThe image %s already exists, thus we do not expect any error", imageName)
		}

	})
	clear()
}

func TestShowDownloadedImages(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if _, err := ShowDownloadedImages(); err != nil {
			t.Errorf("ERROR: on ShowDownloadedImages(): %s", err.Error())
		}
	})
}

func TestImageIsReady(t *testing.T) {
	check := func(t *testing.T, expected bool, result bool) {
		t.Helper()
		if expected != result {
			t.Errorf("Expected %v != Result %v", expected, result)
		}
	}
	clear := func() {
		DeleteImage("alpine")
	}

	// Here we don't call Prepare(), thus we expect that the Image isn't Ready
	Pull("alpine")
	result := ImageIsReady("alpine")
	expected := false
	check(t, expected, result)

	// Here we  call Prepare(), thus we expect that the Image is Ready
	Prepare("alpine", "alpine")
	result = ImageIsReady("alpine")
	expected = true
	check(t, expected, result)

	DeleteImage("alpine")
	result = ImageIsReady("alpine")
	expected = false
	check(t, expected, result)

	clear()
}

func TestDeleteImage(t *testing.T) {
	t.Run("Valid image: busybox", func(t *testing.T) {
		image, containerName := "busybox", "mycontainer"

		if !TarImageExists(image) {
			Pull(image)
		}
		if !ImageIsReady(containerName) {
			Prepare(image, containerName)
		}

		err := DeleteImage(containerName)
		if err != nil {
			t.Errorf(err.Error())
		}
		// certifies that the image was removed
		if ImageIsReady(containerName) {
			t.Errorf("We expected that in this moment the image isn't ready")
		}
	})
	t.Run("Invalid image: busycaixa", func(t *testing.T) {
		containerName := "busycaixa"
		err := DeleteImage(containerName)
		if err != nil {
			t.Errorf("We don't expected any error here, because the image doesn't exists. Err:  %s", err.Error())
		}
	})
	t.Run("Empty image: '' ", func(t *testing.T) {
		image := ""
		err := DeleteImage(image)
		if err.Error() != "The imageName must be a non-empty value" {
			t.Errorf(err.Error())
		}
	})
	clear()
}

func TestDescribe(t *testing.T) {
	t.Run("Unavailable Image", func(t *testing.T) {
		e := ""
		r := Describe("debianxsdsad")
		check(t, e, r)
	})
	t.Run("Alpine Image", func(t *testing.T) {
		e := "alpine\t\tv3.11\t\t2.7M"
		r := Describe("alpine")
		check(t, e, r)
	})
	t.Run("Busybox Image", func(t *testing.T) {
		e := "busybox\t\tv4.0\t\t1.5M"
		r := Describe("busybox")
		check(t, e, r)
	})
	t.Run("Ubuntu Image", func(t *testing.T) {
		e := "ubuntu\t\tv20.04\t\t90.0M"
		r := Describe("ubuntu")
		check(t, e, r)
	})
}

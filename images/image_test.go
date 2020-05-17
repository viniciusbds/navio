package images

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/viniciusbds/navio/utilities"
)

var check = func(t *testing.T, expected string, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %s != Result %s", expected, result)
	}
}

var clear = func() {
	DeleteContImage("alpine")
	DeleteContImage("busybox")
	DeleteContImage("ubuntu")
	os.Remove("images")
}

func TestPull(t *testing.T) {

	//note: this tests don't cover all Pull function

	t.Run("Invalid Image", func(t *testing.T) {
		imageName := "ubuntuxxx"
		err := Pull(imageName)
		if err != nil {
			expected := imageName + " is not a official Image. Select one of the: [alpine busybox ubuntu]"
			result := err.Error()
			check(t, expected, result)
		}
	})

	t.Run("A Image that already was downloaded", func(t *testing.T) {
		imageName := "alpine"

		// first clear a already download image
		os.Remove(filepath.Join(utilities.TarsPath, imageName+".tar"))

		err := Pull(imageName)
		if err != nil {
			t.Errorf("As this is the first pull, is expected a successful pull")
		}

		err = Pull(imageName)
		if err == nil {
			t.Errorf("As this is the second pull, is expected a unsuccessful pull")
		}

		if err != nil {
			expected := "The image " + imageName + " already was downloaded"
			result := err.Error()
			check(t, expected, result)
		}

	})
	clear()
}

func TestShowBaseImages(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if _, err := ShowBaseImages(); err != nil {
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
		DeleteBaseImage("alpine")
	}

	image := "alpine"
	containerImg := "meucontainerzao"

	// Here we don't call Prepare(), thus we expect that the Image isn't Ready
	Pull(image)
	result := IsContImageReady(containerImg)
	expected := false
	check(t, expected, result)

	// Here we  call Prepare(), thus we expect that the Image is Ready on containerImg
	Prepare(image, containerImg)
	result = IsContImageReady(containerImg)
	expected = true
	check(t, expected, result)

	// Here we delete the ready containerImg, thus we expect that is isn't Ready anymore
	DeleteContImage(containerImg)
	result = IsContImageReady(containerImg)
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
		if !IsContImageReady(containerName) {
			Prepare(image, containerName)
		}

		if !IsContImageReady(containerName) {
			t.Errorf("We expected that in this moment the image is ready")
		}

		DeleteContImage(containerName)

		// certifies that the image was removed
		if IsContImageReady(containerName) {
			t.Errorf("We expected that in this moment the image isn't ready")
		}
	})
	clear()
}

package utilities

import "testing"

func TestContains(t *testing.T) {
	imageNames := []string{"alpine", "busybox", "ubuntu", "windows"}
	if !Contains(imageNames, "alpine") || !Contains(imageNames, "busybox") ||
		!Contains(imageNames, "ubuntu") || !Contains(imageNames, "windows") {
		t.Error("Error on Contains function")
	}
	if Contains(imageNames, "kali") || Contains(imageNames, "") {
		t.Error("Error on Contains function")
	}
}

func TestImageisNotEmpty(t *testing.T) {
	t.Run("Simple case", func(t *testing.T) {
		emptyImage := ""
		nonemptyImage := "alpinex"
		if result := IsEmpty(emptyImage); !result {
			t.Errorf("The imageName %s is a empty value", emptyImage)
		}
		if result := IsEmpty(nonemptyImage); result {
			t.Errorf("The imageName %s is a non-empty value", emptyImage)
		}
	})
	t.Run("Empty string with spaces", func(t *testing.T) {
		emptyImage := "      "
		if result := IsEmpty(emptyImage); !result {
			t.Errorf("The imageName %s is a empty value", emptyImage)
		}
	})
}

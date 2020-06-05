package util

import "testing"

func TestImageIsEmpty(t *testing.T) {
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

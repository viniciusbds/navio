package assert

import "testing"

func TestImageisNotEmpty(t *testing.T) {

	t.Run("Simple case", func(t *testing.T) {
		emptyImage := ""
		nonemptyImage := "alpinex"
		if err := ImageisNotEmpty(emptyImage); err.Error() != "The imageName must be a non-empty value" {
			t.Errorf(err.Error())
		}
		if err := ImageisNotEmpty(nonemptyImage); err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("Empty string with spaces", func(t *testing.T) {
		if err := ImageisNotEmpty("    "); err.Error() != "The imageName must be a non-empty value" {
			t.Errorf(err.Error())
		}
	})
}

package assert

import "testing"

func TestImageisNotEmpty(t *testing.T) {
	emptyImage := ""
	nonemptyImage := "alpinex"
	if err := ImageisNotEmpty(emptyImage); err.Error() != "The imageName must be a non-empty value" {
		t.Errorf(err.Error())
	}
	if err := ImageisNotEmpty(nonemptyImage); err != nil {
		t.Errorf(err.Error())
	}
}

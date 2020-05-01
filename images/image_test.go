package images

import (
	"testing"
)

func TestDescribe(t *testing.T) {
	check := func(t *testing.T, expected string, result string) {
		t.Helper()
		if expected != result {
			t.Errorf("Expected %s != Result %s", expected, result)
		}
	}

	t.Run("Unavailable Image", func(t *testing.T) {
		expected := ""
		result := Describe("debianxsdsad")
		check(t, expected, result)
	})

	t.Run("Alpine Image", func(t *testing.T) {
		expected := "alpine\t\tv3.11\t\t2.7M"
		result := Describe("alpine")
		check(t, expected, result)
	})

	t.Run("Busybox Image", func(t *testing.T) {
		expected := "busybox\t\tv4.0\t\t1.5M"
		result := Describe("busybox")
		check(t, expected, result)
	})

	t.Run("Ubuntu Image", func(t *testing.T) {
		expected := "ubuntu\t\tv20.04\t\t90.0M"
		result := Describe("ubuntu")
		check(t, expected, result)
	})
}

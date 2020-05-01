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

package images

import "testing"

func TestToStr(t *testing.T) {

	t.Run("alpine", func(t *testing.T) {
		image := getImage("alpine")
		expected := "alpine\t\tv3.11\t\t2.7M"
		result := image.ToStr()
		check(t, expected, result)
	})

	t.Run("busybox", func(t *testing.T) {
		image := getImage("busybox")
		expected := "busybox\t\tv4.0\t\t1.5M"
		result := image.ToStr()
		check(t, expected, result)
	})

	t.Run("ubuntu", func(t *testing.T) {
		image := getImage("ubuntu")
		expected := "ubuntu\t\tv20.04\t\t90.0M"
		result := image.ToStr()
		check(t, expected, result)
	})

}

func TestGetImage(t *testing.T) {
	t.Run("alpine", func(t *testing.T) {
		image := getImage("alpine")
		if image.name != "alpine" || image.version != "v3.11" || image.size != "2.7M" {
			t.Errorf("Error on GetImage")
		}
	})
}

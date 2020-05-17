package images

import "testing"

func TestGetImage(t *testing.T) {
	baseImg := "alpine"
	t.Run(baseImg, func(t *testing.T) {
		image := getImage(baseImg)
		if image == nil {
			l.Log("INFO", "The "+baseImg+" image  doesn't exists")
			Pull(baseImg)
		}
		image = getImage(baseImg)
		if image == nil {
			t.Error("The " + baseImg + " wasn't download")
		}
		if image.name != baseImg {
			t.Errorf("Error on GetImage")
		}
	})
}

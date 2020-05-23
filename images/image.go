package images

import (
	"fmt"
)

var (
	images = make(map[string]*Image)
)

func init() {
	readImagesDB()
}

// Image holds the structure defining a image object.
type Image struct {
	name    string
	base    string
	version string
	size    string
	url     string
}

// NewImage creates a new image with its basic configuration.
func NewImage(name string, base string, version string, size string, url string) *Image {
	return &Image{
		name:    name,
		base:    base,
		version: version,
		size:    size,
		url:     url,
	}
}

// ToStr ... Ignore esse código, essa foi a minha maior vigarice (https://www.youtube.com/watch?v=PK0c_n5EDhk)
func (i *Image) ToStr() string {
	return fmt.Sprintf("%s            \t\t\t%s\t\t\t%s\t\t\t%s", i.name, i.base, i.version, i.size)
}

func getImage(name string) *Image {
	return images[name]
}

// InsertImage ...
func InsertImage(name, baseImage string) {
	baseImg := getImage(baseImage)
	newImg := NewImage(name, baseImage, baseImg.version, baseImg.size, baseImg.url)
	images[name] = newImg
	insertImageDB(newImg)
}

// IsValid receive a imageName and return true if is a valid image.
func IsValid(image string) bool {
	return getImage(image) != nil
}

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

// Image ...
type Image struct {
	name    string
	base    string
	version string
	size    string
	url     string
}

// NewImage ...
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
	tab := ""
	if len(i.name) < 8 {
		tab = "\t\t\t\t\t"
	} else if len(i.name) >= 8 && len(i.name) < 16 {
		tab = "\t\t\t\t"
	} else {
		tab = "\t\t\t"
	}
	return fmt.Sprintf("%s%s%s\t\t\t%s\t\t\t%s", i.name, tab, i.base, i.version, i.size)
}

func getImage(name string) *Image {
	if images[name] != nil {
		return images[name]
	}
	return nil
}

// InsertImage ...
func InsertImage(name, baseImage string) {
	baseImg := getImage(baseImage)
	newImg := NewImage(name, baseImage, baseImg.version, baseImg.size, baseImg.url)
	images[name] = newImg
	insertImageDB(newImg)
}

// IsaBaseImage receive a imageName and return true if is a base image.
func IsaBaseImage(image string) bool {
	for _, i := range images {
		if image == i.name {
			return true
		}
	}
	return false
}

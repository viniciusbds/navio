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
	Name    string
	Base    string
	Version string
	Size    string
	URL     string
}

// NewImage creates a new image with its basic configuration.
func NewImage(name string, base string, version string, size string, url string) *Image {
	return &Image{
		Name:    name,
		Base:    base,
		Version: version,
		Size:    size,
		URL:     url,
	}
}

// ToStr of the Image
func (i *Image) ToStr() string {
	return fmt.Sprintf("%s            \t\t\t%s\t\t\t%s\t\t\t%s", i.Name, i.Base, i.Version, i.Size)
}

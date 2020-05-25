package images

import (
	"fmt"
	"strings"

	"github.com/viniciusbds/navio/utilities"
)

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
	name := i.Name + strings.Repeat(" ", utilities.MaxImageNameLenght-len(i.Name))
	base := i.Base + strings.Repeat(" ", utilities.MaxImageNameLenght-len(i.Base))
	return fmt.Sprintf("%s%s%s\t\t%s", name, base, i.Version, i.Size)
}

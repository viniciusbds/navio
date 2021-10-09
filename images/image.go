package images

import (
	"fmt"
	"strings"

	"github.com/viniciusbds/navio/constants"
)

// Image holds the structure defining a image object.
type Image struct {
	Name    string
	Base    string
	Version string
	Size    float64
	URL     string
}

// NewImage creates a new image with its basic configuration.
func NewImage(name string, base string, version string, size float64, url string) *Image {
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
	name := i.Name + strings.Repeat(" ", constants.MaxImageNameLength-len(i.Name))
	base := i.Base + strings.Repeat(" ", constants.MaxImageNameLength-len(i.Base))
	return fmt.Sprintf("%s %s %s\t\t%.1f", name, base, i.Version, i.Size)
}

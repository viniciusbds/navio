package images

import "fmt"

var (
	alpine  = NewImage("alpine", "v3.11", "2.7M", "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz")
	busybox = NewImage("busybox", "v4.0", "1.5M", "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar")
	ubuntu  = NewImage("ubuntu", "v20.04", "90.0M", "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz")

	availableImages = []*Image{alpine, busybox, ubuntu}
)

// Image ...
type Image struct {
	name    string
	version string
	size    string
	url     string
}

// NewImage ...
func NewImage(name string, version string, size string, url string) *Image {
	return &Image{
		name:    name,
		version: version,
		size:    size,
		url:     url,
	}
}

// ToStr ...
func (image *Image) ToStr() string {
	return fmt.Sprintf("%s\t\t%s\t\t%s", image.name, image.version, image.size)
}

func getImage(imageName string) *Image {
	for _, i := range availableImages {
		if i.name == imageName {
			return i
		}
	}
	return nil
}

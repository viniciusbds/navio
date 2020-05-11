package images

import (
	"fmt"
)

var (
	// AvailableImages ...
	availableImages = []*Image{
		NewImage("alpine", "alpine", "v3.11", "2.7M", "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz"),
		NewImage("busybox", "busybox", "v4.0", "1.5M", "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar"),
		NewImage("ubuntu", "ubuntu", "v20.04", "90.0M", "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz"),
	}
)

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

// ToStr ...
func (i *Image) ToStr() string {
	return fmt.Sprintf("%s\t\t%s\t\t%s", i.name, i.version, i.size)
}

func getImage(name string) *Image {
	for _, i := range availableImages {
		if i.name == name {
			return i
		}
	}
	return nil
}

// IsValidImage ...
func IsValidImage(name string) bool {
	if getImage(name) != nil {
		return true
	}
	return false
}

// // InsertANewImage ...
// func InsertANewImage(containerName, imageName string, images []*Image) []*Image {
// 	baseImage := getImage(imageName)
// 	image := NewImage(containerName, imageName, baseImage.version, baseImage.size, baseImage.url)
// 	availableImages = append(images, image)
// 	fmt.Printf("insert %v\n", images)
// 	return availableImages
// }

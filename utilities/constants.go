package utilities

// NavioVersion is the current version of Navio
var NavioVersion = "v1.0"

// ImagesRootDir  is the default directory where we manipulate all images
var ImagesRootDir = "/tmp/navioimages"

// ImagesPath is the directory where we store all image.tar files
var ImagesPath = ImagesRootDir + "/tars"

// RootfsPath is the directory where we store all Rootfs directories
var RootfsPath = ImagesRootDir + "/images"

// ContImagescsv is the file where we all container images (for example, if you create a container
// with "ana" as name, the image "ana" will be here)
var ContImagescsv = ImagesRootDir + "/contimages.csv"

// BaseImagescsv ...
var BaseImagescsv = ImagesRootDir + "/baseimages.csv"

// OfficialImages are the official images that are currently suported
var OfficialImages = []string{"alpine", "busybox", "ubuntu"}

// AlpineURL ...
var AlpineURL = "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz"

// BusyboxURL ...
var BusyboxURL = "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar"

// Ubuntu20ltsURL ...
var Ubuntu20ltsURL = "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz"

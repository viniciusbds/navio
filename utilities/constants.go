package utilities

// ImagesRootDir  is the default directory where we manipulate all images
var ImagesRootDir = "/tmp/navioimages"

// TarsPath is the directory where we store all image.tar files
var TarsPath = ImagesRootDir + "/tars"

// ImagesPath is the directory where we store all Images directories
var ImagesPath = ImagesRootDir + "/images"

// Imagescsv is the file where we all container images (for example, if you create a container
// with "ana" as name, the image "ana" will be here)
var Imagescsv = ImagesRootDir + "/imagelist.csv"

// BaseImages are the official images that are currently suported
var BaseImages = []string{"alpine", "busybox", "ubuntu"}

// AlpineURL ...
var AlpineURL = "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz"

// BusyboxURL ...
var BusyboxURL = "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar"

// Ubuntu20ltsURL ...
var Ubuntu20ltsURL = "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz"

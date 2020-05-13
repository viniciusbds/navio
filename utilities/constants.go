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

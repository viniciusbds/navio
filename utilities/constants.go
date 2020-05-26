package utilities

// NavioVersion is the current version of Navio
var NavioVersion = "v1.1"

// RootDir is the place where we store all Navio's data
var RootDir = "/usr/local/navio"

// ImagesPath is the sub-directory where we store all image.tar files.
var ImagesPath = RootDir + "/images"

// RootFSPath is the sub-directory where we store all Rootfs directories.
var RootFSPath = RootDir + "/roots"

// OfficialImages are the official images that are currently suported.
var OfficialImages = []string{"alpine", "busybox", "ubuntu"}

// DBuser represents the database user
var DBuser = "root"

// DBpass represents the database password
var DBpass = "root"

// DBname represents the database name...
var DBname = "navio"

// MaxImageNameLength  23 == len("NAME                    ") -1
var MaxImageNameLength = 23

// MaxContainerNameLength  23 == len("NAME                    ") -1
var MaxContainerNameLength = 23

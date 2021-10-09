package constants

const (
	// NavioVersion is the current version of Navio
	NavioVersion = "v2.0"

	// RootDir is the place where we store all Navio's data
	RootDir = "/usr/local/navio"

	// ImagesPath is the sub-directory where we store all image.tar files.
	ImagesPath = RootDir + "/images"

	// RootFSPath is the sub-directory where we store all Rootfs directories.
	RootFSPath = RootDir + "/roots"

	// DBuser represents the database user
	DBuser = "navioUser"

	// DBpass represents the database password
	DBpass = "PmO001-nav"

	// DBname represents the database name...
	DBname = "navio"

	// MaxImageNameLength  23 == len("NAME                    ") -1
	MaxImageNameLength = 23

	// MaxContainerNameLength  23 == len("NAME                    ") -1
	MaxContainerNameLength = 23

	// DefaultMaxProcessCreation ...
	DefaultMaxProcessCreation = "30"

	// DefaultCPUS ...
	DefaultCPUS = "0"

	// DefaultCPUshares ...
	DefaultCPUshares = "1024"

	// DefaultMemlimit ...
	DefaultMemlimit = "1G"
)

// OfficialImages are the official images that are currently suported.
var OfficialImages = []string{"alpine", "busybox", "ubuntu"}

// IsOfficialImage ...
func IsOfficialImage(image string) bool {
	for _, i := range OfficialImages {
		if image == i {
			return true
		}
	}
	return false
}

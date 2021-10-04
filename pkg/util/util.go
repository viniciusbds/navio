package util

import (
	"os"
	"strings"
	"time"

	"github.com/viniciusbds/navio/pkg/logger"
)

var l = logger.New(time.Kitchen, true)

// Must ....
// [TODO]: Document this function
func Must(err error) {
	if err != nil {
		l.Log("ERROR", err.Error())
		os.Exit(1)
	}
}

// IsEmpty ...
func IsEmpty(imageName string) bool {
	return len(strings.TrimSpace(imageName)) == 0
}

// IsRoot check if the process is running as root user, for example, one started with sudo
func IsRoot() bool {
	return os.Geteuid() == 0
}

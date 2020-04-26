package util

import (
	"os"
	"time"

	"github.com/viniciusbds/navio/src/logger"
)

var l = logger.New(time.Kitchen, true)

// Must ....
func Must(err error) {
	if err != nil {
		l.Log("ERROR", err.Error())
		os.Exit(1)
	}
}

// Contains ...
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

package assert

import (
	"errors"
	"time"

	"github.com/viniciusbds/navio/logger"
)

var l = logger.New(time.Kitchen, true)

// ImageisNotEmpty ...
func ImageisNotEmpty(imageName string) error {
	if imageName == "" {
		err := errors.New("The imageName must be a non-empty value")
		l.Log("ERROR", err.Error())
		return err
	}
	return nil
}

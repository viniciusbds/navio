package assert

import (
	"errors"
	"strings"
	"time"

	"github.com/viniciusbds/navio/logger"
)

var l = logger.New(time.Kitchen, true)

// ImageisNotEmpty ...
func ImageisNotEmpty(imageName string) error {
	if len(strings.TrimSpace(imageName)) == 0 {
		err := errors.New("The imageName must be a non-empty value")
		l.Log("ERROR", err.Error())
		return err
	}
	return nil
}

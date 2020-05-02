package logger

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	l := New(time.Kitchen, true)
	l.Log("INFO", "This is a info message")
	l.Log("wARNING", "This is a warning message")
	l.Log("ERROR", "This is a error message")
}

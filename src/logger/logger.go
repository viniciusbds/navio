package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

var red = ansi.ColorFunc("red+")
var green = ansi.ColorFunc("green+")
var magenta = ansi.ColorFunc("magenta+")

// Logger ...
type Logger struct {
	timeFormat string
	debug      bool
}

// New ...
// [TODO]: Document the constructor
func New(timeFormat string, debug bool) *Logger {
	return &Logger{timeFormat: timeFormat, debug: debug}
}

// Log ...
// [TODO]: Document this function
func (l *Logger) Log(level, message string) {
	switch strings.ToLower(level) {
	case "info":
		if l.debug {
			l.print(level, message, green)
		}
	case "warning":
		if l.debug {
			l.print(level, message, magenta)
		}
	default:
		l.print(level, message, red)
	}
}

func (l *Logger) print(level string, message string, collorfunc func(string) string) {
	fmt.Printf(collorfunc("[%s] %s --> %s\n"), level, time.Now().Format(l.timeFormat), message)
}

package naviofile

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/viniciusbds/navio/pkg/logger"
)

var (
	imgTag string
	l      = logger.New(time.Kitchen, true)
)

// ReadNaviofile ...
func ReadNaviofile(path string) (baseImage, source, destination string, commands [][]string) {

	naviofile, err := ioutil.ReadFile(filepath.Join(path, "Naviofile")) // just pass the file name
	if err != nil {
		l.Log("ERROR", err.Error())
		return
	}

	lines := strings.Split(string(naviofile), "\n")
	for _, line := range lines {
		l := strings.Split(line, " ")
		cmd := l[0]

		if cmd == "FROM" {
			baseImage = l[1]
		} else if cmd == "ADD" {
			source = l[1]
			destination = l[2]
		} else if cmd == "RUN" {
			l = strings.Split(line, "&&")
			// expected example: [RUN apt update,  apt -y upgrade, apt install -y python]

			for i, c := range l {
				c = strings.TrimSpace(c)
				aux := strings.Split(c, " ")
				if i == 0 {
					// removing the the [RUN] cmd
					aux = aux[1:]
				}

				commands = append(commands, aux)
			}
		}
	}
	return baseImage, source, destination, commands
}

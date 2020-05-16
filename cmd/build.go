package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/utilities"
)

func init() {
	rootCmd.AddCommand(build())
}

func build() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Create a new Image",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			// navio build [image-name] [directory]

			var naviofileDir, newImg string
			if len(args) < 2 {
				l.Log("WARNING", "You must insert a directory and a image name!")
				return
			}

			if utilities.FileExists(args[0] + "/Naviofile") {
				naviofileDir = args[0]
				newImg = args[1]
			} else if utilities.FileExists(args[1] + "/Naviofile") {
				naviofileDir = args[1]
				newImg = args[0]
			} else {
				l.Log("WARNING", "You must insert a directory of your Naviofile!")
				return
			}

			naviofile, err := ioutil.ReadFile(filepath.Join(naviofileDir, "Naviofile")) // just pass the file name
			if err != nil {
				l.Log("ERROR", err.Error())
				return
			}

			var from, origem, destino string
			var commands = [][]string{}

			lines := strings.Split(string(naviofile), "\n")
			for _, line := range lines {
				l := strings.Split(line, " ")
				cmd := l[0]

				if cmd == "FROM" {
					from = l[1]
				} else if cmd == "ADD" {
					origem = l[1]
					destino = l[2]
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

			fmt.Printf("FROM %s\n", from)
			fmt.Printf("ADD %s %s\n", origem, destino)
			fmt.Printf("RUN %v\n", commands)
			fmt.Println("------------------")

			// FROM
			images.BuildANewBaseImg(newImg, from)

			// ADD
			// [TODO]

			// RUN
			for _, command := range commands {
				container.CreateContainer(append([]string{from, newImg}, command...))
			}

			// saving the image.tarin tarPath ...
			dir := filepath.Join(utilities.ImagesPath, newImg)
			file := filepath.Join(utilities.TarsPath, newImg+".tar")
			utilities.Must(utilities.Tar(dir, file))

			images.InsertBaseImage(newImg, from)

		},
	}
}

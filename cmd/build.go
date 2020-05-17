package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/utilities"
)

var (
	newImageName string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&newImageName, "t", "", "The image tag. (i.e. the newImageName)")
	rootCmd.MarkFlagRequired("t")

	rootCmd.AddCommand(build())
}

func build() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Create a new Image",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			// navio build [directory] -t [image-name]

			if utilities.IsEmpty(newImageName) {
				l.Log("WARNING", "You must insert a image name. for ex.: --t python-ubuntu")
				return
			}

			if len(args) < 1 {
				l.Log("WARNING", "You must insert a directory of your Naviofile!")
				return
			}
			if len(args) > 1 {
				l.Log("WARNING", "You only need insert a directory of your Naviofile!")
				return
			}
			if !utilities.FileExists(args[0] + "/Naviofile") {
				l.Log("WARNING", "You must insert a directory of your Naviofile!")
				return
			}

			naviofileDir := args[0]

			naviofile, err := ioutil.ReadFile(filepath.Join(naviofileDir, "Naviofile")) // just pass the file name
			if err != nil {
				l.Log("ERROR", err.Error())
				return
			}

			var baseImage, origem, destino string
			var commands = [][]string{}

			lines := strings.Split(string(naviofile), "\n")
			for _, line := range lines {
				l := strings.Split(line, " ")
				cmd := l[0]

				if cmd == "FROM" {
					baseImage = l[1]
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

			fmt.Printf("FROM %s\n", baseImage)
			fmt.Printf("ADD %s %s\n", origem, destino)
			fmt.Printf("RUN %v\n", commands)
			fmt.Println("------------------")

			// FROM
			images.BuildANewBaseImg(newImageName, baseImage)

			// ADD
			// [TODO]

			// RUN
			for _, command := range commands {
				container.CreateContainer(append([]string{baseImage, newImageName}, command...))
			}

			// saving the image.tarin tarPath ...
			dir := filepath.Join(utilities.ImagesPath, newImageName)
			file := filepath.Join(utilities.TarsPath, newImageName+".tar")
			utilities.Must(utilities.Tar(dir, file))

			images.InsertBaseImage(newImageName, baseImage)

			// clear the rootfs used to build the image.tar file
			if err := os.RemoveAll(filepath.Join(utilities.ImagesPath, newImageName)); err != nil {
				l.Log("ERROR", err.Error())
				return
			}

		},
	}
}

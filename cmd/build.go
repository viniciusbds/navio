package cmd

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/naviofile"
	"github.com/viniciusbds/navio/pkg/loader"
	"github.com/viniciusbds/navio/utilities"
)

var (
	imgTag string
	wg     sync.WaitGroup
	done   chan bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&imgTag, "t", "", "The image tag. (i.e. the newImageName)")
	rootCmd.MarkFlagRequired("t")
	rootCmd.AddCommand(build())

	done = make(chan bool)
}

func build() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build an image from a Naviofile",
		Run: func(cmd *cobra.Command, args []string) {

			if utilities.IsEmpty(imgTag) {
				l.Log("WARNING", "You must insert a image name. for ex.: --t python-ubuntu")
				return
			}
			if len(imgTag) > utilities.MaxImageNameLength {
				l.Log("WARNING", "Image name is too long, please enter a shorter name.")
				return
			}
			if images.Exists(imgTag) {
				l.Log("WARNING", "This image name already exists.")
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
			baseImage, source, destination, commands := naviofile.ReadNaviofile(naviofileDir)

			fmt.Printf(magenta("FROM %s\n"), baseImage)
			fmt.Printf(magenta("ADD %s %s\n"), source, destination)
			fmt.Printf(magenta("RUN %v\n"), commands)
			fmt.Printf("---------------------------------------------------------------\n")

			rand.Seed(time.Now().UnixNano())
			min := 11111111
			max := 99999999
			containerID := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)
			containerRootFS := filepath.Join(utilities.RootFSPath, containerID)

			// FROM
			fmt.Printf(green("Copying the [%s] image ...\n"), baseImage)
			wg.Add(1)
			go loader.Loader("Done :)", done, &wg)
			go images.UntarImg(baseImage, containerRootFS, done)
			wg.Wait()

			// ADD
			if !utilities.IsEmpty(source) && !utilities.IsEmpty(destination) {
				fmt.Printf(green("ADD %s %s\n"), source, destination)
				fullDestinyPath := filepath.Join(containerRootFS, destination)
				wg.Add(1)
				go loader.Loader("Done :)", done, &wg)
				go utilities.Copy(source, fullDestinyPath, done)
				wg.Wait()
			}

			// ENTRYPOINT
			// [TODO]

			// ENV
			// [TODO]

			// WORKDIR
			// [TODO]

			// CMD
			// [TODO]

			// RUN
			containerName := imgTag
			command := "echo"
			params := []string{"Creating", "this", "container", "just", "to", "run", "the", "commands", "to", "build", "a", "new", "image"}
			go container.CreateContainer(containerID, containerName, baseImage, command, params, done)

			fmt.Printf(green("Prepare container ...\n"))
			wg.Add(1)
			go loader.Loader("Done :)", done, &wg)
			wg.Wait()

			for _, c := range commands {
				command := c[0]
				params := c[1:]
				fmt.Printf(green("RUN %v\n"), append([]string{command}, params...))
				container.Exec(containerID, containerName, command, params)
			}

			// saving the image.tarin tarPath ...
			imageFile := filepath.Join(utilities.ImagesPath, imgTag+".tar")

			fmt.Printf(green("Generating the [%s] image ...\n"), imgTag)
			wg.Add(1)
			go loader.Loader("Done :)", done, &wg)
			go utilities.Tar(containerRootFS, imageFile, done)
			wg.Wait()

			images.InsertImage(imgTag, baseImage)
			err := container.RemoveContainer(containerID)
			if err != nil {
				l.Log("ERROR", err.Error())
			}
		},
	}
}

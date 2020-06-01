package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/naviofile"
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

			// navio build [directory] -t [image-name]

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
			baseImage, source, destiny, commands := naviofile.ReadNaviofile(naviofileDir)

			fmt.Printf(magenta("FROM %s\n"), baseImage)
			fmt.Printf(magenta("ADD %s %s\n"), source, destiny)
			fmt.Printf(magenta("RUN %v\n"), commands)
			fmt.Printf("---------------------------------------------------------------\n")

			containerID := fmt.Sprintf("%d", rand.Int31n(1000000000))
			containerName := imgTag

			if container.RootfsExists(containerName) {
				l.Log("WARNING", fmt.Sprintf("The containerName %s already was used. Enter a new name.", containerName))
				os.Exit(1)
			}

			containerRootFS := filepath.Join(utilities.RootFSPath, containerName)

			// FROM
			fmt.Printf(green("Copying the [%s] image ...\n"), baseImage)
			go images.UntarImg(imgTag, baseImage, done)

			// ADD
			if !utilities.IsEmpty(source) && !utilities.IsEmpty(destiny) {
				fullDestinyPath := filepath.Join(containerRootFS, destiny)
				wg.Add(1)
				go utilities.Copy(source, fullDestinyPath, &wg)
			}

			// Wait the copy of the IMAGE complete
			wg.Add(1)
			go utilities.Loader(done, &wg)
			wg.Wait()

			// ENTRYPOINT
			// [TODO]

			// ENV
			// [TODO]

			// WORKDIR
			// [TODO]

			// CMD
			// [TODO]

			// RUN
			args = append([]string{baseImage, containerID, containerName, "echo"}, []string{"Creating", "this", "container", "just", "to", "run", "the", "commands", "to", "build", "a", "new", "image"}...)
			go container.CreateContainer(args, done)

			fmt.Printf(green("Prepare container ...\n"))
			wg.Add(1)
			go utilities.Loader(done, &wg)
			wg.Wait()

			for _, command := range commands {
				fmt.Printf(green("RUN %v\n"), command)
				container.Exec(append([]string{containerName}, command...))
			}

			// saving the image.tarin tarPath ...
			imageFile := filepath.Join(utilities.ImagesPath, imgTag+".tar")

			fmt.Printf(green("Generating the [%s] image ...\n"), imgTag)
			wg.Add(1)
			go utilities.Loader(done, &wg)
			go utilities.Tar(containerRootFS, imageFile, done)
			wg.Wait()

			images.InsertImage(imgTag, baseImage)
			err := container.RemoveContainer(containerName)
			if err != nil {
				l.Log("ERROR", err.Error())
			}
		},
	}
}

package cmd

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/constants"
	"github.com/viniciusbds/navio/containers"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/naviofile"
	"github.com/viniciusbds/navio/pkg/io"
	"github.com/viniciusbds/navio/pkg/spinner"
	"github.com/viniciusbds/navio/pkg/util"
)

var (
	imgTag string
	wg     sync.WaitGroup
	done   chan bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&imgTag, "t", "", "The image tag. (i.e. the newImageName)")
	err := rootCmd.MarkFlagRequired("t")
	if err != nil {
		l.Log("ERROR", err.Error())
	}

	rootCmd.AddCommand(build())

	done = make(chan bool)
}

func build() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build an image from a Naviofile",
		Run: func(cmd *cobra.Command, args []string) {

			if !util.IsRoot() {
				l.Log("WARNING", "This command requires sudo privileges! please run as super user :)")
				return
			}

			if util.IsEmpty(imgTag) {
				l.Log("WARNING", "You must insert a image name. for ex.: --t python-ubuntu")
				return
			}
			if len(imgTag) > constants.MaxImageNameLength {
				l.Log("WARNING", "Image name is too long, please enter a shorter name.")
				return
			}
			if images.IsAvailable(imgTag) {
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
			if !io.FileExists(args[0] + "/Naviofile") {
				l.Log("WARNING", "You must insert a directory of your Naviofile!")
				return
			}

			naviofileDir := args[0]
			baseImage, source, destination, commands := naviofile.ReadNaviofile(naviofileDir)

			fmt.Printf(magenta("FROM %s\n"), baseImage)
			fmt.Printf(magenta("ADD %s %s\n"), source, destination)
			fmt.Printf(magenta("RUN %v\n"), commands)
			fmt.Printf("---------------------------------------------------------------\n")

			containerID := containers.GenerateNewID()
			containerRootFS := filepath.Join(constants.RootFSPath, containerID)

			// FROM
			fmt.Printf(green("Copying the [%s] image ...\n"), baseImage)
			wg.Add(1)
			go spinner.Spinner("Done :)", done, &wg)
			errs := make(chan error, 1)
			go func() {
				errs <- images.Untar(baseImage, containerRootFS, done)
			}()
			wg.Wait()
			if err := <-errs; err != nil {
				l.Log("ERROR", err.Error())
			}

			// ADD
			if !util.IsEmpty(source) && !util.IsEmpty(destination) {
				fmt.Printf(green("ADD %s %s\n"), source, destination)
				fullDestinyPath := filepath.Join(containerRootFS, destination)
				wg.Add(1)
				go spinner.Spinner("Done :)", done, &wg)
				go func() {
					errs <- io.Copy(source, fullDestinyPath, done)
				}()
				wg.Wait()
				if err := <-errs; err != nil {
					l.Log("ERROR", err.Error())
				}

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

			cgroups := containers.NewCGroup(pids, cpus, cpushares, memory)
			go func() {
				errs <- containers.CreateContainer(containerID, containerName, baseImage, command, params, done, cgroups)
			}()
			fmt.Print(green("Prepare container	 ...\n"))
			wg.Add(1)
			go spinner.Spinner("Done :)", done, &wg)
			wg.Wait()
			if err := <-errs; err != nil {
				l.Log("ERROR", err.Error())
			}

			for _, c := range commands {
				command := c[0]
				params := c[1:]
				fmt.Printf(green("RUN %v\n"), append([]string{command}, params...))
				err := containers.Exec(containerID, command, params)
				if err != nil {
					l.Log("ERROR", err.Error())
				}

			}

			// saving the image.tarin tarPath ...
			imageFile := filepath.Join(constants.ImagesPath, imgTag+".tar")

			fmt.Printf(green("Generating the [%s] image ...\n"), imgTag)
			wg.Add(1)
			go spinner.Spinner("Done :)", done, &wg)
			go func() {
				errs <- io.Tar(containerRootFS, imageFile, done)
			}()
			wg.Wait()
			if err := <-errs; err != nil {
				l.Log("ERROR", err.Error())
			}

			imageSize, err := io.FileSize(imageFile)
			if err != nil {
				l.Log("Error in image size calculation", err.Error())
			}
			imageSizeInMB := float64(imageSize) / 1000000

			err = images.Insert(imgTag, imageSizeInMB, baseImage)
			if err != nil {
				l.Log("Error on insert image", err.Error())
			}

			err = containers.Remove(containerID)
			if err != nil {
				l.Log("ERROR", err.Error())
			}
		},
	}
}

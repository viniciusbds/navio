package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/constants"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/pkg/spinner"
	"github.com/viniciusbds/navio/pkg/util"
)

var (
	// Used for name flag.
	containerName string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&containerName, "name", "", "The name of the container")
	rootCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new container",
		Run: func(cmd *cobra.Command, args []string) {

			if !util.IsRoot() {
				l.Log("WARNING", "This command requires sudo privileges! please run as super user :)")
				return
			}

			if len(containerName) > constants.MaxContainerNameLength {
				l.Log("WARNING", "Container name is too long, please enter a shorter name.")
				return
			}

			image, index := getImage(args)
			if image == "" {
				l.Log("WARNING", "Insert a valid image name.")
				return
			}

			// remove the image of args
			args = append(args[:index], args[index+1:]...)
			if len(args) == 0 {
				l.Log("WARNING", "You must insert a command.")
				return
			}
			command, params := args[0], args[1:]

			containerID := container.GenerateNewID()

			if containerName == "" {
				containerName = containerID
			}

			if container.UsedName(containerName) {
				l.Log("WARNING", fmt.Sprintf("The containerName %s was already used. Enter a new name.", containerName))
				os.Exit(1)
			}

			fmt.Printf(green("Image: %s, Command: %s, Params: %v\n"), image, command, params)

			fmt.Printf(green("Creating [%s] container ...\n"), containerName)
			wg.Add(1)
			go spinner.Spinner("Done :)", done, &wg)
			container.CreateContainer(containerID, containerName, image, command, params, done)
		},
	}
}

func getImage(args []string) (image string, index int) {
	var arg string
	for index, arg = range args {
		if images.GetImage(arg) != nil {
			image = arg
			break
		}
	}
	return
}

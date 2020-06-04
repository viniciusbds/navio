package cmd

import (
	"fmt"
	"os"
	"time"

	"math/rand"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/utilities"
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
		Use:   "run",
		Short: "Run a command in a new container",
		Run: func(cmd *cobra.Command, args []string) {

			if len(containerName) > utilities.MaxContainerNameLength {
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

			rand.Seed(time.Now().UnixNano())
			min := 99999999
			max := 1000000000
			containerID := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)

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
			go utilities.Loader(done, &wg)
			container.CreateContainer(containerID, containerName, image, command, params, done)

		},
	}
}

func getImage(args []string) (image string, index int) {
	var arg string
	for index, arg = range args {
		if images.IsValid(arg) {
			image = arg
			break
		}
	}
	return
}

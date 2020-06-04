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
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio run IMAGE COMMAND PARAMS...

			if len(args) < 2 {
				l.Log("WARNING", "You must insert at least a image name and a command!")
				return nil
			}

			if len(containerName) > utilities.MaxContainerNameLength {
				l.Log("WARNING", "Container name is too long, please enter a shorter name.")
				return nil
			}

			image := args[0]

			if !images.IsValid(image) {
				l.Log("WARNING", fmt.Sprintf("%s is not a base Image. See navio get images", image))
				return nil
			}

			command := args[1]
			params := args[2:]

			rand.Seed(time.Now().UnixNano())
			min := 99999999
			max := 1000000000
			containerID := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)

			if containerName == "" {
				containerName = containerID
			}

			if container.RootFSExists(containerID) {
				l.Log("WARNING", fmt.Sprintf("The containerName %s already was used. Enter a new name.", containerName))
				os.Exit(1)
			}

			fmt.Printf(green("Image: %s, Command: %s, Params: %v\n"), image, command, params)

			fmt.Printf(green("Creating [%s] container ...\n"), containerName)
			wg.Add(1)
			go utilities.Loader(done, &wg)
			container.CreateContainer(containerID, containerName, image, command, params, done)

			return nil
		},
	}
}

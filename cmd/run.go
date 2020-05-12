package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
)

var (
	// Used for containerame flag.
	containerName string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&containerName, "name", "", "The name of the container")
	rootCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio run IMAGE COMMAND PARAMS...

			if len(args) < 2 {
				l.Log("WARNING", "You must insert at least a image name and a command!")
				return nil
			}

			image := args[0]

			if !images.IsValidImage(image) {
				l.Log("WARNING", "Invalid Image!!! To see all available images run navio get images")
				return nil
			}

			command := args[1]
			params := args[2:]

			if containerName == "" {
				// TODO: generate a random container name
				containerName = "XPTO"
			}

			l.Log("INFO", fmt.Sprintf("Image: %s, Command: %s, Params: %v", image, command, params))
			args = append([]string{image, command, containerName}, params...)
			container.CreateContainer(args)

			return nil
		},
	}
}

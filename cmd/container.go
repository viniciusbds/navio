package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
)

func init() {
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio run IMAGE COMMAND PARAMS...
			image := args[0]
			command := args[1]
			params := args[2:]

			l.Log("INFO", fmt.Sprintf("Image: %s, Command: %s, Params: %v", image, command, params))

			args = append([]string{"run", image, command}, params...)
			container.CreateContainer(args)

			return nil
		},
	}
}

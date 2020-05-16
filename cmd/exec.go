package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(exec())
}

func exec() *cobra.Command {
	return &cobra.Command{
		Use: "exec",
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio exec CONTAINERIMAGENAME COMMAND PARAMS...

			if len(args) < 2 {
				l.Log("WARNING", "You must insert at least a containerName and a command!")
				return nil
			}

			containerName := args[0]

			if !images.IsValidContainerImage(containerName) {
				l.Log("WARNING", fmt.Sprintf("%s is not a valid containerImage. Run navio ps to see the available ones.", containerName))
				return nil
			}
			command := args[1]
			params := args[2:]

			l.Log("INFO", fmt.Sprintf("Image: %s, Command: %s, Params: %v", containerName, command, params))
			args = append([]string{containerName, containerName, command}, params...)
			container.CreateContainer(args)

			return nil
		},
	}
}

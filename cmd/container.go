package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/src/container"
)

func init() {
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Args: ", args)

			// navio run IMAGE COMMAND PARAMS...
			image := args[0]
			command := args[1]
			params := args[2:]

			args = append([]string{"runa", image, command}, params...)

			container.CreateContainer(args)

			return nil
		},
	}
}

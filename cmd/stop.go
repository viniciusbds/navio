package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(stop())
}

func stop() *cobra.Command {
	return &cobra.Command{
		Use: "stop",
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio stop CONTAINERIMAGENAME

			if len(args) == 0 {
				l.Log("WARNING", "You must insert the containerName!")
				return nil
			}
			if len(args) > 1 {
				l.Log("WARNING", "You only need insert the containerName!")
				return nil
			}

			containerName := args[0]

			if !images.IsValidContainerImage(containerName) {
				l.Log("WARNING", fmt.Sprintf("%s is not a valid containerImage. Run navio ps to see the available ones.", containerName))
				return nil
			}

			images.DeleteImage(containerName)

			return nil
		},
	}
}

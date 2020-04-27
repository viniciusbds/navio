package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(pullImage())
}

func pullImage() *cobra.Command {
	return &cobra.Command{
		Use: "pull",
		RunE: func(cmd *cobra.Command, args []string) error {
			image := args[0]
			images.Pull(image)
			return nil
		},
	}
}

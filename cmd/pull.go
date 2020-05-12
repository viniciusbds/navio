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
		Use:   "pull",
		Short: "Download an image from its official website.",
		Long: "Receive an [image_name] as an argument and download it from the official website." +
			" To see all available images run: navio get images",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				l.Log("WARNING", "You must insert a image name!")
			} else {
				image := args[0]
				images.Pull(image)
			}

			return nil
		},
	}
}

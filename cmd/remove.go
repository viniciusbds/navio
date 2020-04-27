package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(remove())
}

func remove() *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "Remove Navio objects",
		Long:  "ex: navio remove image <image_name> remove a downloaded images located in the ./images directory.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if args[0] == "image" {
				// testar caso não exista arg[1]
				images.RemoveDownloadedImages(args[1])
			}

			return nil
		},
	}
}

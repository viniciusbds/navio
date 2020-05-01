package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(describe())
}

func describe() *cobra.Command {
	return &cobra.Command{
		Use:   "describe",
		Short: "Describe Navio objects",
		Long:  "ex: navio describe image <image_name> describe available image to download.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if args[0] == "image" {
				// testar caso não exista arg[1]
				fmt.Println(images.Describe(args[1]))
			}

			return nil
		},
	}
}

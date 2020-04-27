package cmd

import (
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

var magenta = ansi.ColorFunc("magenta+")

func init() {
	rootCmd.AddCommand(get())
}

func get() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "....",
		Long:  "....",
		RunE: func(cmd *cobra.Command, args []string) error {

			if args[0] == "images" {
				showDownloadedImages()
			}

			return nil
		},
	}
}

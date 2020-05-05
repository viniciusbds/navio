package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/utilities"
)

func init() {
	rootCmd.AddCommand(get())
}

func get() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Return Navio objects",
		Long:  "i.e: navio get images show all downloaded images that are in the ./images directory.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				l.Log("WARNING", "Insert a valid argument ex: images")
				return nil
			}

			if args[0] == "images" {
				fmt.Println("NAME\t\tVERSION\t\tSIZE")
				imageList, _ := images.ShowDownloadedImages()
				if !utilities.IsEmpty(imageList) {
					fmt.Println(imageList)
				}
			}

			return nil
		},
	}
}

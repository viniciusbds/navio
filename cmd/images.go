package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(listImages())
}

func listImages() *cobra.Command {
	return &cobra.Command{
		Use:   "images",
		Short: "List images",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("NAME\t\t\tBASE\t\t\tVERSION\t\tSIZE\t" + images.ListImages())
		},
	}
}

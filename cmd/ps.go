package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
	"github.com/viniciusbds/navio/utilities"
)

func init() {
	rootCmd.AddCommand(ps())
}

func ps() *cobra.Command {
	return &cobra.Command{
		Use:   "ps",
		Short: "Shows all containerImages that was created",
		Long:  "Each of thesees containerImages are a /rootf of the respective container that was created",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("NAME\t\t\t\t\tBASE\t\t\tVERSION\t\t\tSIZE")
			imageList, _ := images.Ps()
			if !utilities.IsEmpty(imageList) {
				fmt.Println(imageList)
			}

			return nil
		},
	}
}

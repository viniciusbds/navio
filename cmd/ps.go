package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
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
			fmt.Println("ID\t\tNAME\t\t\tIMAGE\t\t\tCOMMAND\t\t\tSTATUS")
			containers, _ := container.Ps()
			if !utilities.IsEmpty(containers) {
				fmt.Println(containers)
			}
			return nil
		},
	}
}

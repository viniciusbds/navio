package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
)

func init() {
	rootCmd.AddCommand(containers())
}

func containers() *cobra.Command {
	return &cobra.Command{
		Use:   "containers",
		Short: "List containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ID\t\tNAME\t\t\tIMAGE\t\t\tCOMMAND\t\t\tSTATUS")
			list, _ := container.List()
			fmt.Println(list)
		},
	}
}

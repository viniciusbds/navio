package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/containers"
)

func init() {
	rootCmd.AddCommand(listContainers())
}

func listContainers() *cobra.Command {
	return &cobra.Command{
		Use:   "containers",
		Short: "List all containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ID\t\tNAME\t\t\tIMAGE\t\t\tCOMMAND\t\t\tSTATUS")
			list, _ := containers.List()
			fmt.Println(list)
		},
	}
}

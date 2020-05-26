package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rm())
}

func rm() *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove one or more containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("|----- TODO -----|")

		},
	}
}

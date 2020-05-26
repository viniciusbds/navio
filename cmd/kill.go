package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// |----- TODO -----|

func init() {
	rootCmd.AddCommand(kill())
}

func kill() *cobra.Command {
	return &cobra.Command{
		Use:   "kill",
		Short: "Kill one or more running containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("|----- TODO -----|")
		},
	}
}

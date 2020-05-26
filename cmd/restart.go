package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// |----- TODO -----|

func init() {
	rootCmd.AddCommand(restart())
}

func restart() *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "Restart one or more containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("|----- TODO -----|")
		},
	}
}

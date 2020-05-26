package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// |----- TODO -----|

func init() {
	rootCmd.AddCommand(rename())
}

func rename() *cobra.Command {
	return &cobra.Command{
		Use:   "rename",
		Short: "Rename a container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("|----- TODO -----|")
		},
	}
}

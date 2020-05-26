package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// |----- TODO -----|

func init() {
	rootCmd.AddCommand(unpause())
}

func unpause() *cobra.Command {
	return &cobra.Command{
		Use:   "unpause",
		Short: "Unpause all processes within one or more containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("|----- TODO -----|")
		},
	}
}

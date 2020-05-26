package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// |----- TODO -----|

func init() {
	rootCmd.AddCommand(pause())
}

func pause() *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "Pause all processes within one or more containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("|----- TODO -----|")
		},
	}
}

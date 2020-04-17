package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Navio",
	Long:  "All software has versions. This is Navio's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Navio 0.1")
	},
}

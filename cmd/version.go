package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Navio version",
	Long:  "All software has versions. This is Navio's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("|___/ Navio ", version)
	},
}

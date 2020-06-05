package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/constants"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Navio version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("|___/ Navio ", constants.NavioVersion)
	},
}

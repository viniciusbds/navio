package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/utilities"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Navio version",
	Long:  "All software has versions. This is Navio's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("|___/ Navio ", utilities.NavioVersion)
	},
}

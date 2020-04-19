package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "navio",
		Short:        "Atenção tripulação, os containers estão surgindo!",
		SilenceUsage: true,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("Root cmd")
		// },
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

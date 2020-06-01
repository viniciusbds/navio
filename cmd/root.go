package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/logger"
)

var (
	l       = logger.New(time.Kitchen, true)
	red     = ansi.ColorFunc("red+")
	green   = ansi.ColorFunc("green+")
	magenta = ansi.ColorFunc("magenta+")
)

var (
	rootCmd = &cobra.Command{
		Use:          "navio",
		Short:        "|___/ Navio is an extremely simple app that creates linux containers",
		SilenceUsage: true,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("Root cmd")
		// },
	}
)

// Execute executes the root command.
// [TODO]: Document this function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	//"github.com/viniciusbds/navio/src" m
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/viniciusbds/navio"
)

func init() {
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args)
			if contains(args, "ubuntu") {
				// do ...
			}
			//m.cria()
			return nil
		},
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/src"
)

func init() {
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use: "container",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Args: ", args)

			// navio run [IMAGE, "child"] COMMAND PARAMS...
			command := args[1]
			params := args[2:]

			// if args[0] == "arch" {
			// 	// do...
			// }

			// if args[0] == "alpine" {
			// 	// do...
			// }

			// if args[0] == "ubuntu" {
			// 	// do...
			// }

			if args[0] == "child" {
				args = append([]string{"child"}, command)
				args = append(args, params...)
			} else {
				args = append([]string{"run"}, command)
				args = append(args, params...)
			}
			src.CreateContainer(args)

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

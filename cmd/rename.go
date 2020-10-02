package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/containers"
	"github.com/viniciusbds/navio/constants"
)

var (
	id string
	name string
)

func init() {
	rootCmd.AddCommand(rename())
}

func rename() *cobra.Command {
	return &cobra.Command{
		Use:   "rename",
		Short: "Rename a container",
		Run: func(cmd *cobra.Command, args []string) {
			id, name = args[0], args[1]
			if !containers.IsaID(id) {
				fmt.Println(red("ERROR: Container not exists"))
			} else if len(name) > constants.MaxContainerNameLength {
				fmt.Println(red("ERROR: Container name is too long, please enter a shorter name."))
			} else {
				containers.UpdateName(id, name)
				fmt.Println(green("Container renamed successfully!"))
			}
		},
	}
}
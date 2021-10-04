package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/containers"
	"github.com/viniciusbds/navio/pkg/util"
)

func init() {
	rootCmd.AddCommand(exec())
}

func exec() *cobra.Command {
	return &cobra.Command{
		Use:   "exec",
		Short: "Run a command in a container",
		Run: func(cmd *cobra.Command, args []string) {

			if !util.IsRoot() {
				l.Log("WARNING", "This command requires sudo privileges! please run as super user :)")
				return
			}

			containerID, indexID := getContainerID(args)
			if containerID == "" {
				l.Log("WARNING", "Insert a valid container id.")
				return
			}

			// remove the containerID of args
			args = append(args[:indexID], args[indexID+1:]...)
			if len(args) == 0 {
				l.Log("WARNING", "You must insert a command.")
				return
			}

			command, params := args[0], args[1:]
			l.Log("INFO", fmt.Sprintf("Container: %s, Command: %s, Params: %v", containers.GetContainerName(containerID), command, params))
			err := containers.Exec(containerID, command, params)
			if err != nil {
				l.Log("ERROR", err.Error())
			}
		},
	}
}

func getContainerID(args []string) (containerID string, index int) {
	var arg string
	for index, arg = range args {
		if containers.IsaID(arg) {
			containerID = arg
			break
		}
	}
	return
}

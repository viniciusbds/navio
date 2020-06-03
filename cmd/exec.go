package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/utilities"
)

var (
	// Used for id flag.
	containerID string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&containerID, "id", "", "The ID of the container")
	rootCmd.AddCommand(exec())
}

func exec() *cobra.Command {
	return &cobra.Command{
		Use:   "exec",
		Short: "Run a command in a running container",
		Run: func(cmd *cobra.Command, args []string) {
			if utilities.IsEmpty(containerID) {
				l.Log("WARNING", "You must insert a container id. (ex --id 9999987)")
				return
			}
			if !utilities.IsEmpty(containerID) && !container.IsaID(containerID) {
				l.Log("WARNING", "Invalid container id. Run [navio containers] to see the available ones.")
				return
			}
			if len(args) < 1 {
				l.Log("WARNING", "You must insert a command!")
				return
			}

			command := args[0]
			params := args[1:]

			l.Log("INFO", fmt.Sprintf("Container: %s, Command: %s, Params: %v", containerName, command, params))
			args = append([]string{containerID, containerName, command}, params...)
			container.Exec(args)
		},
	}
}

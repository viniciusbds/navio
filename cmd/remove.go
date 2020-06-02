package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
)

func init() {
	rootCmd.AddCommand(remove())
}

func remove() *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "Remove one or more containers",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				l.Log("WARNING", "You must insert the containerName!")
				return
			}

			for _, arg := range args {

				if !container.Exists(arg) {
					l.Log("WARNING", "The container "+arg+" doesn't exists!")
					continue
				}

				if !container.IsID(arg) {
					arg = container.GetContainerID(arg)
				}

				err := container.RemoveContainer(arg)
				if err != nil {
					l.Log("ERROR", err.Error())
				}
			}

		},
	}
}

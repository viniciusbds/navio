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

			if args[0] == "all" {
				err := container.RemoveAll()
				if err != nil {
					l.Log("ERROR", err.Error())
				}
			} else {
				var id string
				for _, arg := range args {

					if !container.IsaID(arg) && !container.UsedName(arg) {
						l.Log("WARNING", "The container "+arg+" doesn't exists!")
						continue
					}

					if container.IsaID(arg) {
						id = arg
					} else {
						id = container.GetContainerID(arg)
					}

					if err := container.RemoveContainer(id); err != nil {
						l.Log("ERROR", err.Error())
					}
				}
			}

		},
	}
}

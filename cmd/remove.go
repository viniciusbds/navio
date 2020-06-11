package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/containers"
	"github.com/viniciusbds/navio/pkg/spinner"
	"github.com/viniciusbds/navio/pkg/util"
)

func init() {
	rootCmd.AddCommand(remove())
}

func remove() *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "Remove one or more containers",
		Run: func(cmd *cobra.Command, args []string) {

			if !util.IsRoot() {
				l.Log("WARNING", "This command requires sudo privileges! please run as super user :)")
				return
			}

			if len(args) == 0 {
				l.Log("WARNING", "You must insert the containerName!")
				return
			}

			if args[0] == "all" {
				wg.Add(1)
				fmt.Println("Removing all containers ...")
				go containers.RemoveAll(done)
				spinner.Spinner("Done :)", done, &wg)
				wg.Wait()
			} else {
				var id string
				for _, arg := range args {

					if !containers.IsaID(arg) && !containers.UsedName(arg) {
						l.Log("WARNING", "The container "+arg+" doesn't exists!")
						continue
					}

					if containers.IsaID(arg) {
						id = arg
					} else {
						id = containers.GetContainerID(arg)
					}

					if err := containers.RemoveContainer(id); err != nil {
						l.Log("ERROR", err.Error())
					}
				}
			}

		},
	}
}

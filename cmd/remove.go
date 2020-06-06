package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/isroot"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/pkg/spinner"
)

func init() {
	rootCmd.AddCommand(remove())
}

func remove() *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "Remove one or more containers",
		Run: func(cmd *cobra.Command, args []string) {

			if !isroot.IsRoot() {
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
				go container.RemoveAll(done)
				spinner.Spinner("Done :)", done, &wg)
				wg.Wait()
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

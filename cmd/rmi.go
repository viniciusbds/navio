package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusbds/isroot"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(rmi())
}

func rmi() *cobra.Command {
	return &cobra.Command{
		Use:   "rmi",
		Short: "Remove a image",
		Long:  "ex: navio remove image <image_name> remove a downloaded images located in the ./images directory.",
		Run: func(cmd *cobra.Command, args []string) {

			if !isroot.IsRoot() {
				l.Log("WARNING", "This command requires sudo privileges! please run as super user :)")
				return
			}

			if len(args) == 0 {
				l.Log("WARNING", "You must insert at least a image name!")
				return
			}

			if args[0] == "all" {
				err := images.RemoveAll()
				if err != nil {
					l.Log("ERROR", err.Error())
				}
			} else {
				for _, arg := range args {
					if arg != "" {
						err := images.Remove(arg)
						if err != nil {
							l.Log("ERROR", err.Error())
						}
					}
				}
			}
		},
	}
}

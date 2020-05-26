package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
)

func init() {
	rootCmd.AddCommand(stop())
}

func stop() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop one or more running containers",
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio stop CONTAINERIMAGENAME

			if len(args) == 0 {
				l.Log("WARNING", "You must insert the containerName!")
				return nil
			}

			for _, containerName := range args {
				container.RemoveContainer(containerName)
			}

			return nil
		},
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(pullImage())
}

func pullImage() *cobra.Command {
	return &cobra.Command{
		Use:   "pull",
		Short: "Download an image from its official website.",
		Long: "Receive an [image_name] as an argument and download it from the official website." +
			" To see all available images run: navio get images",
		RunE: func(cmd *cobra.Command, args []string) error {
			image := args[0]
			images.Pull(image)
			return nil
		},
	}
}

func showDownloadedImages() {
	dirname := "./images"

	f, err := os.Open(dirname)
	if err != nil {
		l.Log("ERROR", err.Error())
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		l.Log("ERROR", err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(magenta(file.Name()))
		}
	}
}

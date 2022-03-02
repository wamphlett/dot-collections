package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wamphlett/dot-collections/pkg/application"
)

//go:embed templates/collections.sh
var collectionsFile embed.FS

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs dot collections for the first time",

	Run: func(cmd *cobra.Command, args []string) {
		// check that it is not already installed
		if application.IsInstalled() {
			fmt.Println("already installed!")
			return
		}

		application.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

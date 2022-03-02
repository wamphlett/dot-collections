package cmd

import (
	"fmt"
	"github.com/wamphlett/dot-collections/pkg/application"
	"os"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "removes all collections and uninstalls dot-collections",

	Run: func(cmd *cobra.Command, args []string) {
		if err := os.RemoveAll(application.InstallLocation()); err != nil {
			fmt.Printf("failed to remove install files: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("dot-collections was successfully purged!")
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}

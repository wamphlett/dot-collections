package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the getCollection command
var getCmd = &cobra.Command{
	Use:   "get [collection]",
	Short: "Display collection details",

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()
		collection, err := app.Get(args[0])
		if err != nil {
			fmt.Printf("failed to get collection deatils: %s\n", err.Error())
			os.Exit(1)
		}

		printVariables(collection, true)
		printVariables(collection, false)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCollectionsCmd represents the listCollections command
var listCollectionsCmd = &cobra.Command{
	Use:   "list",
	Short: "List the installed collections",
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()

		allCollections, err := app.GetAll()
		if err != nil {
			fmt.Printf("failed to load collections: %s\n", err.Error())
			os.Exit(1)
		}

		for _, collection := range allCollections {
			desc := collection.Description
			if desc == "" {
				desc = "no description"
			}
			fmt.Printf("%s: %s\n", collection.Slug, desc)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCollectionsCmd)
}

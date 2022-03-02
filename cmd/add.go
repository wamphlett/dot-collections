package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wamphlett/dotc/pkg/application"
	"os"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Clones an existing collection repository",

	Run: func(cmd *cobra.Command, args []string) {
		var url, slug string
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the git address for the collection:")
		if scanner.Scan() {
			url = scanner.Text()
		}
		fmt.Println("Enter an identifiable name for the collection (no spaces):")
		if scanner.Scan() {
			slug = scanner.Text()
		}
		application.CloneCollection(slug, url)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

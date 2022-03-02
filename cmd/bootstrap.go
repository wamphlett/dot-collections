package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap [collection]",
	Short: "Run the collections bootstrap file",
	Long:  `Runs the collections specified bootstrap file as defined in the collection.yaml file.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()

		collection, err := app.Get(args[0])
		if err != nil {
			fmt.Printf("failed to load collection: %s\n", err.Error())
			os.Exit(1)
		}

		bootstrapPath := filepath.Join(collection.Path, collection.Bootstrapper)
		if collection.Bootstrapper == "" {
			fmt.Println("No bootstrap file specified.")
			return
		}

		// check the bootstrap file exists
		if _, err := os.Stat(bootstrapPath); os.IsNotExist(err) {
			fmt.Printf("Specified bootstrap file does not exist: %s", collection.Bootstrapper)
			return
		}

		// execute the bootstrap command
		c := exec.Command("sh", bootstrapPath)

		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		c.Start()
		c.Wait()
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}

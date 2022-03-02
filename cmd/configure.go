/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/wamphlett/dotc/pkg/core/collections"
	"os"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure [collection]",
	Short: "Configure collection environment variables",

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()
		collection, err := app.Get(args[0])
		if err != nil {
			fmt.Printf("failed to load collection: %s\n", err.Error())
			os.Exit(1)
		}

		for _, variable := range collection.Variables {
			configureVariable(variable)
		}

		fmt.Println("New values:")
		printVariables(collection, true)
		printVariables(collection, false)

		var commit string
		for commit != "y" && commit != "n" {
			fmt.Println("would you like to save changes? (y/n):")
			fmt.Scanln(&commit)
		}

		if commit == "n" {
			fmt.Println("your changes have been discarded")
			return
		}

		if err := app.UpdateVariables(collection); err != nil {
			fmt.Printf("failed to update settings: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println("your changes have been saved")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func configureVariable(variable *collections.Variable) {
	fmt.Printf("Key: %s\n", variable.Key)
	fmt.Printf("Default value: %s\n", variable.Default)
	fmt.Printf("Current value: %s\n", variable.Value)

	if variable.Value == "" && variable.Default != "" {
		fmt.Println("no current value. use default? (y)")
		var useDefault string
		fmt.Scanln(&useDefault)
		if useDefault == "y" {
			variable.Value = variable.Default
			return
		}
	}

	fmt.Println("New value (leave blank to keep current value):")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		newValue := scanner.Text()
		if newValue != "" {
			variable.Value = newValue
		}
	}
}

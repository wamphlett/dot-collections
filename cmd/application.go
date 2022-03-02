package cmd

import (
	"fmt"
	"github.com/wamphlett/dot-collections/pkg/application"
	"github.com/wamphlett/dot-collections/pkg/configurator"
	"github.com/wamphlett/dot-collections/pkg/core/collections"
)

func getApp() *collections.Service {
	application.EnsureInstalled()
	installPath := application.InstallLocation()
	return collections.New(installPath, configurator.New(installPath))
}

func printVariables(collection *collections.Collection, exported bool) {
	if exported == true {
		fmt.Println("Environment variables:")
	} else {
		fmt.Println("Shell variables:")
	}

	for _, variable := range collection.Variables {
		if variable.IsEnv != exported {
			continue
		}
		fmt.Printf(" - %s: %s\n", variable.Key, variable.Value)
	}
}

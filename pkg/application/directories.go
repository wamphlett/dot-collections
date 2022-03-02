package application

import (
	"fmt"
	"os"
	"path"
)

func InstallLocation() string {
	return path.Join(HomeDir(), ".dot-collections")
}

func HomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("failed to locate home directory: %s\n", err.Error())
		os.Exit(1)
	}
	return dir
}

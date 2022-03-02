package application

import (
	"embed"
	"fmt"
	"os"
	"path"
)

//go:embed templates/*
var collectionsFile embed.FS

func IsInstalled() bool {
	if _, err := os.Stat(InstallLocation()); os.IsNotExist(err) {
		return false
	}
	return true
}

func EnsureInstalled() {
	if IsInstalled() == false {
		fmt.Println("It looks like this is the first time you are using dot-collections. Running install...")
		Install()
	}
}

func Install() {
	// make the install directory
	fmt.Println("creating install folder...")
	if err := os.MkdirAll(InstallLocation(), 0755); err != nil {
		fmt.Printf("failed to create install directory: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("copying install files...")
	data, _ := collectionsFile.ReadFile("templates/collections.sh")
	if err := os.WriteFile(path.Join(InstallLocation(), "collections.sh"), data, 0644); err != nil {
		fmt.Printf("failed to create install files: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Install successful.\n\nadd \"source %s/collections.sh\" to you bashrc file\n\nRun \"dotc add\" to add your first collection\n\n",
		InstallLocation())
}

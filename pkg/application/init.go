package application

import (
	"os"
	"os/exec"
	"path/filepath"
)

func CloneCollection(slug, url string) {
	installPath := filepath.Join(InstallLocation(), slug)
	// execute the bootstrap command
	c := exec.Command("git", "clone", url, installPath)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Start()
	c.Wait()
}

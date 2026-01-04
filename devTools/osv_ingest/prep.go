package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const dataScript = "osv_prep.sh"

func runPrepScript() error {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	// Check if script exists in current directory
	scriptPath := filepath.Join(wd, dataScript)
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script not found: %s in %s", dataScript, wd)
	}

	fmt.Printf("Running script: %s\n", scriptPath)

	// Execute the bash script
	// Use relative path for bash compatibility on Windows
	cmd := exec.Command("bash", dataScript)
	cmd.Dir = wd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute script: %w", err)
	}

	fmt.Println("Script executed successfully")
	return nil
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// RunDecrypt executes the decryption and setup commands
func RunDecrypt() {
	fmt.Println("Starting decryption and configuration setup...")

	// Command 1: decrypt the gpg file
	fmt.Println("Running: gpg rclone.conf.gpg")
	cmdDecrypt := exec.Command("gpg", "rclone.conf.gpg")
	cmdDecrypt.Stdout = os.Stdout
	cmdDecrypt.Stderr = os.Stderr
	err := cmdDecrypt.Run()
	if err != nil {
		fmt.Printf("Error decrypting file: %v\n", err)
		return
	}

	// Command 2: remove the gpg file
	fmt.Println("Running: rm rclone.conf.gpg")
	cmdRemove := exec.Command("rm", "rclone.conf.gpg")
	cmdRemove.Stdout = os.Stdout
	cmdRemove.Stderr = os.Stderr
	err = cmdRemove.Run()
	if err != nil {
		fmt.Printf("Error removing gpg file: %v\n", err)
		// Continue even if there's an error
	}

	// Command 3: move the conf file to ~/.config/rclone/
	// First ensure the directory exists
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	rcloneConfigDir := filepath.Join(homeDir, ".config", "rclone")

	// Create directory if it doesn't exist
	err = os.MkdirAll(rcloneConfigDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", rcloneConfigDir, err)
		return
	}

	fmt.Printf("Moving rclone.conf to %s\n", rcloneConfigDir)
	cmdMove := exec.Command("mv", "rclone.conf", filepath.Join(rcloneConfigDir))
	cmdMove.Stdout = os.Stdout
	cmdMove.Stderr = os.Stderr
	err = cmdMove.Run()
	if err != nil {
		fmt.Printf("Error moving config file: %v\n", err)
		return
	}

	fmt.Println("Decrypt operations completed successfully.")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if any command-line arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main [command]")
		fmt.Println("Available commands:")
		fmt.Println("  init     - Install dependencies and download configuration")
		fmt.Println("  decrypt  - Decrypt and install rclone configuration")
		return
	}

	// Get the command from arguments
	command := os.Args[1]

	// Execute the appropriate command based on the argument
	switch command {
	case "init":
		RunInit()
	case "decrypt":
		RunDecrypt()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands:")
		fmt.Println("  init     - Install dependencies and download configuration")
		fmt.Println("  decrypt  - Decrypt and install rclone configuration")
	}
}

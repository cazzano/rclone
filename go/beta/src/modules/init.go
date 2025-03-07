package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Starting apt operations...")

	// Command 1: apt upgrade
	fmt.Println("Running: apt upgrade")
	cmdUpgrade := exec.Command("apt", "upgrade")
	cmdUpgrade.Stdout = os.Stdout
	cmdUpgrade.Stderr = os.Stderr
	err := cmdUpgrade.Run()
	if err != nil {
		fmt.Printf("Error running apt upgrade: %v\n", err)
	}

	// Command 2: apt install gnupg -y
	fmt.Println("Running: apt install gnupg -y")
	cmdInstallGnupg := exec.Command("apt", "install", "gnupg", "-y")
	cmdInstallGnupg.Stdout = os.Stdout
	cmdInstallGnupg.Stderr = os.Stderr
	err = cmdInstallGnupg.Run()
	if err != nil {
		fmt.Printf("Error installing gnupg: %v\n", err)
	}

	// Command 3: apt install wget
	fmt.Println("Running: apt install wget")
	cmdInstallWget := exec.Command("apt", "install", "wget")
	cmdInstallWget.Stdout = os.Stdout
	cmdInstallWget.Stderr = os.Stderr
	err = cmdInstallWget.Run()
	if err != nil {
		fmt.Printf("Error installing wget: %v\n", err)
	}

	// Command 4: wget to download rclone.conf.gpg
	fmt.Println("Running: wget https://github.com/cazzano/rclone/raw/main/secrets/rclone.conf.gpg")
	cmdWget := exec.Command("wget", "https://github.com/cazzano/rclone/raw/main/secrets/rclone.conf.gpg")
	cmdWget.Stdout = os.Stdout
	cmdWget.Stderr = os.Stderr
	err = cmdWget.Run()
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
	}

	fmt.Println("All operations completed.")
}

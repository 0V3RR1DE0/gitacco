package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: gitacco add <username> <email>")
			return
		}
		username := os.Args[2]
		email := os.Args[3]
		addAccount(username, email)
	case "switch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gitacco switch <username>")
			return
		}
		username := os.Args[2]
		switchAccount(username)
	case "list":
		listAccounts()
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gitacco remove <username>")
			return
		}
		username := os.Args[2]
		removeAccount(username)
	case "current":
		showCurrentAccount()
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("GitAcco - Git Account Manager")
	fmt.Println("\nUsage:")
	fmt.Println("  gitacco add <username> <email>       - Add a new account")
	fmt.Println("  gitacco switch <username>            - Switch to an account")
	fmt.Println("  gitacco list                         - List all accounts")
	fmt.Println("  gitacco remove <username>            - Remove an account")
	fmt.Println("  gitacco current                      - Show current account")
}
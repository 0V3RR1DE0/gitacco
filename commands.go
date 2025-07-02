package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func addAccount(username, email string) {
	config := loadConfig()
	
	account := Account{
		Username: username,
		Email:    email,
	}
	
	config.Accounts[username] = account
	saveConfig(config)
	
	fmt.Printf("✅ Account '%s' added successfully!\n", username)
	fmt.Printf("   Username: %s\n", username)
	fmt.Printf("   Email: %s\n", email)
}

func switchAccount(username string) {
	config := loadConfig()
	
	account, exists := config.Accounts[username]
	if !exists {
		fmt.Printf("❌ Account '%s' not found. Use 'gitacco list' to see available accounts.\n", username)
		return
	}

	// Set global git config
	err := runGitCommand("config", "--global", "user.name", account.Username)
	if err != nil {
		fmt.Printf("❌ Error setting git user.name: %v\n", err)
		return
	}

	err = runGitCommand("config", "--global", "user.email", account.Email)
	if err != nil {
		fmt.Printf("❌ Error setting git user.email: %v\n", err)
		return
	}

	fmt.Printf("✅ Switched to account '%s'\n", username)
	fmt.Printf("   Username: %s\n", account.Username)
	fmt.Printf("   Email: %s\n", account.Email)
}

func listAccounts() {
	config := loadConfig()
	
	if len(config.Accounts) == 0 {
		fmt.Println("No accounts configured. Use 'gitacco add' to add an account.")
		return
	}

	fmt.Println("📋 Configured accounts:")
	for username, account := range config.Accounts {
		fmt.Printf("  • %s (%s)\n", username, account.Email)
	}
}

func removeAccount(username string) {
	config := loadConfig()
	
	if _, exists := config.Accounts[username]; !exists {
		fmt.Printf("❌ Account '%s' not found.\n", username)
		return
	}

	delete(config.Accounts, username)
	saveConfig(config)
	
	fmt.Printf("✅ Account '%s' removed successfully!\n", username)
}

func showCurrentAccount() {
	name, err := runGitCommandOutput("config", "--global", "user.name")
	if err != nil {
		fmt.Printf("❌ Error getting current git user.name: %v\n", err)
		return
	}

	email, err := runGitCommandOutput("config", "--global", "user.email")
	if err != nil {
		fmt.Printf("❌ Error getting current git user.email: %v\n", err)
		return
	}

	fmt.Println("🔍 Current Git configuration:")
	fmt.Printf("   Username: %s\n", strings.TrimSpace(name))
	fmt.Printf("   Email: %s\n", strings.TrimSpace(email))
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	return cmd.Run()
}

func runGitCommandOutput(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	return string(output), err
}
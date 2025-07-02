package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func addAccount(username, email string) {
	config := loadConfig()
	
	account := Account{
		Username:       username,
		Email:          email,
		GitHubUsername: username, // Use same username for GitHub
	}
	
	config.Accounts[username] = account
	saveConfig(config)
	
	fmt.Printf("✅ Account '%s' added successfully!\n", username)
	fmt.Printf("   Username: %s\n", username)
	fmt.Printf("   Email: %s\n", email)
	fmt.Printf("   GitHub Username: %s\n", username)
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

	// Set GitHub credential (using the same username)
	err = runGitCommand("config", "--global", "credential.https://github.com.username", account.GitHubUsername)
	if err != nil {
		fmt.Printf("⚠️  Warning: Could not set GitHub credential: %v\n", err)
	} else {
		fmt.Printf("🔑 GitHub credential set for: %s\n", account.GitHubUsername)
	}

	fmt.Printf("✅ Switched to account '%s'\n", username)
	fmt.Printf("   Username: %s\n", account.Username)
	fmt.Printf("   Email: %s\n", account.Email)
	fmt.Printf("   GitHub Username: %s\n", account.GitHubUsername)
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

	githubUser, err := runGitCommandOutput("config", "--global", "credential.https://github.com.username")
	
	fmt.Println("🔍 Current Git configuration:")
	fmt.Printf("   Username: %s\n", strings.TrimSpace(name))
	fmt.Printf("   Email: %s\n", strings.TrimSpace(email))
	if err == nil && strings.TrimSpace(githubUser) != "" {
		fmt.Printf("   GitHub Username: %s\n", strings.TrimSpace(githubUser))
	}
}

func setGitHubCredential(username string) {
	config := loadConfig()
	
	account, exists := config.Accounts[username]
	if !exists {
		fmt.Printf("❌ Account '%s' not found. Add it first with 'gitacco add'.\n", username)
		return
	}

	err := runGitCommand("config", "--global", "credential.https://github.com.username", account.GitHubUsername)
	if err != nil {
		fmt.Printf("❌ Error setting GitHub credential: %v\n", err)
		return
	}

	fmt.Printf("✅ GitHub credential set for account '%s'\n", username)
	fmt.Printf("   GitHub Username: %s\n", account.GitHubUsername)
	fmt.Println("💡 Git will now remember to use this account for GitHub operations.")
}

func unsetGitHubCredential() {
	err := runGitCommand("config", "--global", "--unset", "credential.https://github.com.username")
	if err != nil {
		fmt.Printf("❌ Error unsetting GitHub credential: %v\n", err)
		return
	}

	fmt.Println("✅ GitHub credential removed")
	fmt.Println("💡 Git will now prompt for credentials on GitHub operations.")
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
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Account struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	GitHubUsername string `json:"github_username,omitempty"` // Optional GitHub username
}

type Config struct {
	Accounts map[string]Account `json:"accounts"`
}

func getConfigPath() string {
	// Get the directory where the executable is located
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		os.Exit(1)
	}
	
	execDir := filepath.Dir(execPath)
	return filepath.Join(execDir, "gitacco-config.json")
}

func loadConfig() *Config {
	configPath := getConfigPath()
	
	// Create config file if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := &Config{
			Accounts: make(map[string]Account),
		}
		saveConfig(config)
		return config
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}

	if config.Accounts == nil {
		config.Accounts = make(map[string]Account)
	}

	return &config
}

func saveConfig(config *Config) {
	configPath := getConfigPath()
	
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling config: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
		os.Exit(1)
	}
}

func showConfigLocation() {
	fmt.Printf("📁 Config file location: %s\n", getConfigPath())
}
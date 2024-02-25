// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"encoding/json"
	"os"
	"path"
)

const (
	// We will check this environment variable to find the mai configuration directory.
	maiConfigDirEnvVar = "MAI_CONFIG_DIR"
	// If the environment variable is not set, we will use this directory in the user's home directory.
	defaultConfigDir = ".mai"
	// The name of the configuration file.
	configFileName = "config.json"
)

// Represents a single model in the configuration.
type Model struct {
	Name         string            `json:"name"`
	Provider     string            `json:"provider"`
	ProviderName string            `json:"providerName"`
	Args         map[string]string `json:"args"`
}

// Represents the configuration of the mai inatallaion.
type AppConfig struct {
	Models       []*Model `json:"models"`
	DefaultModel string   `json:"defaultModel"`
}

// Infer location of the config file.
func GetAppConfigLocation() (string, error) {
	configDir := os.Getenv(maiConfigDirEnvVar)
	if configDir == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		configDir = path.Join(homedir, defaultConfigDir)
	}

	return path.Join(configDir, configFileName), nil
}

// Check if the configuration file exists.
func AppConfigExists() (bool, error) {
	configFilePath, err := GetAppConfigLocation()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Parse and get the configuration of this mai installation.
func GetAppConfig() (*AppConfig, error) {
	configFilePath, err := GetAppConfigLocation()
	if err != nil {
		return nil, err
	}

	c, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	err = json.Unmarshal(c, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}

// Creates a new empty configuration file and replaces the existing one if necessary.
func InitAppConfig() error {
	configFilePath, err := GetAppConfigLocation()
	if err != nil {
		return err
	}

	// Create the file with all directory parents if it doesn't exist
	err = os.MkdirAll(path.Dir(configFilePath), 0755)
	if err != nil {
		return err
	}

	// Create the file or replace it if it exists
	c, err := json.MarshalIndent(&AppConfig{
		Models: []*Model{},
	}, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, c, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Deletes the configuration file.
func NukeAppConfig() error {
	configFilePath, err := GetAppConfigLocation()
	if err != nil {
		return err
	}

	err = os.Remove(configFilePath)
	if err != nil {
		return err
	}

	return nil
}

// Write the configuration to the file.
func WriteAppConfig(appConfig *AppConfig) error {
	configFilePath, err := GetAppConfigLocation()
	if err != nil {
		return err
	}

	c, err := json.MarshalIndent(appConfig, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, c, 0644)
	if err != nil {
		return err
	}

	return nil
}

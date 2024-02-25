// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"fmt"
	"os"

	"github.com/satmihir/mai/config"
	"github.com/spf13/cobra"
)

// configAddModelCmd represents the configAddModel command
var ConfigAddModelCmd = &cobra.Command{
	Use:   "add-model",
	Short: "Add a model to the mai configuration",
	Long:  `Add a model to the mai configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.GetAppConfig()
		if err != nil {
			fmt.Printf("Failed to get config: %s\n", err)
			os.Exit(1)
		}

		fmt.Print("Select your model provider (openai): ")
		var provider string
		fmt.Scanln(&provider)

		// Just for now, we only support openai.
		if provider != "openai" {
			fmt.Println("Invalid provider.")
			os.Exit(1)
		}

		fmt.Print("Select the name of the provider's model (gpt-3.5-turbo, gpt-4): ")
		var providerName string
		fmt.Scanln(&providerName)

		// Just for now, we only support gpt-3.5-turbo and gpt-4.
		if providerName != "gpt-3.5-turbo" && providerName != "gpt-4" {
			fmt.Println("Invalid model name.")
			os.Exit(1)
		}

		fmt.Print("Choose a name for your model: ")
		var modelName string
		fmt.Scanln(&modelName)

		// Check if the model name is valid.
		if modelName == "" {
			fmt.Println("Invalid model name.")
			os.Exit(1)
		}

		// Check if the model name already exists.
		for _, model := range conf.Models {
			if model.Name == modelName {
				fmt.Println("Model name already exists.")
				os.Exit(1)
			}
		}

		fmt.Print("Enter the openai API key for accessing this model: ")
		var apiKey string
		fmt.Scanln(&apiKey)

		// Check if the API key is valid.
		if apiKey == "" {
			fmt.Println("Invalid API key.")
			os.Exit(1)
		}

		// Add the model to the configuration.
		conf.Models = append(conf.Models, &config.Model{
			Name:         modelName,
			Provider:     provider,
			ProviderName: providerName,
			Args: map[string]string{
				"apiKey": apiKey,
			},
		})

		// TODO: Try a prompt before writing the config.

		err = config.WriteAppConfig(conf)
		if err != nil {
			fmt.Printf("Failed to write config: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Added the model to the configuration.")
	},
}

func init() {}

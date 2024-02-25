// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"fmt"
	"os"

	"github.com/satmihir/mai/config"
	"github.com/spf13/cobra"
)

// configSetDefaultCmd represents the configSetDefault command
var ConfigSetDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set a model as default",
	Long:  `Set a model as default. If there's no explicit default, the first model will be set as default.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := cmd.Flag("model").Value.String()

		conf, err := config.GetAppConfig()
		if err != nil {
			fmt.Printf("Failed to get config: %s\n", err)
			os.Exit(1)
		}

		// Check if the model exists.
		var modelExists bool
		for _, m := range conf.Models {
			if m.Name == model {
				modelExists = true
				break
			}
		}

		if !modelExists {
			fmt.Println("Model does not exist.")
			os.Exit(1)
		}

		conf.DefaultModel = model

		err = config.WriteAppConfig(conf)
		if err != nil {
			fmt.Printf("Failed to save config: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	ConfigSetDefaultCmd.PersistentFlags().StringP("model", "m", "", "The name of the model to set as default.")
	ConfigSetDefaultCmd.MarkPersistentFlagRequired("model")
}

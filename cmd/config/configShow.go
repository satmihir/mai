// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/satmihir/mai/config"
	"github.com/spf13/cobra"
)

// configShowCmd represents the configShow command
var ConfigShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Pretty print the config for mai",
	Long:  `Pretty print the config for mai`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.GetAppConfig()
		if err != nil {
			fmt.Println("Failed to get config.")
			os.Exit(1)
		}

		confJ, err := json.MarshalIndent(conf, "", "  ")
		if err != nil {
			fmt.Println("Failed to marshal config.")
			os.Exit(1)
		}

		fmt.Println(string(confJ))
	},
}

func init() {}

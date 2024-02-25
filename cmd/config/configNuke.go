// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"fmt"
	"os"

	"github.com/satmihir/mai/config"
	"github.com/spf13/cobra"
)

// configNukeCmd represents the configNuke command
var ConfigNukeCmd = &cobra.Command{
	Use:   "nuke",
	Short: "Nuke the mai config file",
	Long:  `Nuke the mai config file`,
	Run: func(cmd *cobra.Command, args []string) {
		exists, err := config.AppConfigExists()
		if err != nil {
			fmt.Printf("Failed to check if config file exists: %s\n", err)
			os.Exit(1)
		}

		if !exists {
			fmt.Println("Config file does not exist.")
			os.Exit(0)
		} else {
			err = config.NukeAppConfig()
			if err != nil {
				fmt.Printf("Failed to nuke config file: %s\n", err)
				os.Exit(1)
			}

			fmt.Println("Nuked the config file.")
		}
	},
}

func init() {}

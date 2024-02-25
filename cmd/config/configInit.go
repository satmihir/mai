// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"fmt"
	"os"

	"github.com/satmihir/mai/config"
	"github.com/spf13/cobra"
)

// configInitCmd represents the configInit command
var ConfigInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize mai configuration",
	Long:  `Initialize mai configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var fileLoc string

		fileLoc, err = config.GetAppConfigLocation()
		if err != nil {
			fmt.Printf("Failed to get config file location: %s\n", err)
			os.Exit(1)
		}

		exists, err := config.AppConfigExists()
		if err != nil {
			fmt.Printf("Failed to check if config file exists: %s\n", err)
			os.Exit(1)
		}

		if exists {
			fmt.Println("Config file already exists. If you want to recreate, delete the existing with `mai config nuke`.")
			os.Exit(0)
		} else {
			err = config.InitAppConfig()
			if err != nil {
				fmt.Printf("Failed to initialize config file: %s\n", err)
				os.Exit(1)
			}

			fmt.Printf("Wrote config file to: %s\n", fileLoc)
		}
	},
}

func init() {}

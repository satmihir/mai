// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the CLI",
	Long:  `Configure the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config base command is a no-op. Run `mai config --help for options`")
	},
}

func init() {
}

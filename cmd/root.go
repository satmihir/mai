// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"os"

	"github.com/satmihir/mai/cmd/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mai",
	Short: "A commandline tool for conversational AI.",
	Long:  `A command line tool for conversational AI.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(config.ConfigCmd)

	// Subcommands of config
	config.ConfigCmd.AddCommand(config.ConfigInitCmd)
	config.ConfigCmd.AddCommand(config.ConfigNukeCmd)
	config.ConfigCmd.AddCommand(config.ConfigAddModelCmd)
	config.ConfigCmd.AddCommand(config.ConfigShowCmd)
	config.ConfigCmd.AddCommand(config.ConfigSetDefaultCmd)
}

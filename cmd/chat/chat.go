// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package chat

import (
	"fmt"
	"os"

	"github.com/satmihir/mai/config"
	"github.com/satmihir/mai/llm"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Send a chat request",
	Long:  `Send a chat request`,
	Run: func(cmd *cobra.Command, args []string) {
		promptString := cmd.Flag("inline").Value

		conf, err := config.GetAppConfig()
		if err != nil {
			fmt.Println("Failed to get app config")
			os.Exit(1)
		}

		mr, err := llm.NewModelRegistry(conf)
		if err != nil {
			fmt.Println("Failed to get model registry")
			os.Exit(1)
		}

		model, ok := mr.GetDefaultModel()
		if !ok {
			fmt.Println("Failed to get the default model")
			os.Exit(1)
		}

		hist := make([]*llm.Message, 0)
		hist = append(hist, &llm.Message{
			Role:    "user",
			Content: promptString.String(),
			Name:    "user",
		})

		completion, err := model.Complete(&llm.Prompt{
			MessgeHistory: hist,
			Choices:       1,
		})

		if err != nil {
			fmt.Printf("Failed invoking the model [%s]\n", err)
			os.Exit(1)
		}

		choice := completion.Choices[0]
		fmt.Print(choice.Text)
	},
}

func init() {
	ChatCmd.PersistentFlags().StringP("inline", "i", "", "Run the chat inline using default model.")
	ChatCmd.MarkPersistentFlagRequired("inline")
}

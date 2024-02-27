// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package llm

import (
	"io"

	"github.com/pkg/errors"
	"github.com/satmihir/mai/config"
)

// Represents a single message in the message history sent to the model.
type Message struct {
	// Role of the message sender such as "system" or "user".
	// Actual exhaustive list of roles is defined by the model.
	Role string
	// The content (text only) of the message.
	Content string
	// Optional name of the actor to discriminate between multiple users
	// in the same conversation.
	Name string
}

// Represents the information sent to the model in a prompt.
// The model name and hyper params are not parts of this struct.
type Prompt struct {
	// A history of messages sent to the model, oldest first.
	// The latest message is the one the model immediately responds to.
	MessgeHistory []*Message
	// The maximum number of tokens the model is allowed to generate.
	MaxTokens uint32
	// Number of choices to generate for each prompt.
	Choices uint32
	// A boolean flag to indicate the answer should be streamed.
	// Choices must be 1 in this case.
	Stream bool
}

// Statistics about the usage of the model while answering a prompt.
type UsageStats struct {
	// Number of tokens in prompt text.
	PromptTokens uint32
	// Number of tokens in completed text.
	CompletionTokens uint32
	// Total number of tokens used in a given prompt.
	TotalTokens uint32
}

// Represents a single choice in the completion.
type CompletionChoice struct {
	// The reason the text finished.
	FinishReason string
	// The index of the choice in the list of choices.
	Index uint32
	// The message that was generated.
	Text string
	// In case of a streaming completion, a stream of text.
	TextStream io.Reader
}

// Represents the completion of a prompt.
type Completion struct {
	// Stats of tokens used in the completion.
	UsageStats *UsageStats
	// The choices of the completion.
	Choices []*CompletionChoice
}

// Represents an LLM model interface
type LLM interface {
	// Completes a prompt and returns the completion.
	Complete(prompt *Prompt) (*Completion, error)
}

// The registry to hold and vend models for the app
type ModelRegistry struct {
	registry         map[string]LLM
	defaultModelName string
}

// Build a model registry for the given AppConfig
func NewModelRegistry(appConfig *config.AppConfig) (*ModelRegistry, error) {
	var defaultModelName string
	reg := make(map[string]LLM)
	for i, m := range appConfig.Models {
		if i == 0 {
			defaultModelName = m.Name
		}

		switch m.Provider {
		case "openai":
			l, err := NewOpenAiLlm(m)
			if err != nil {
				return nil, err
			}

			reg[m.Name] = l
		default:
			return nil, errors.Errorf("Provider %s not supported", m.Provider)
		}
	}

	return &ModelRegistry{
		registry:         reg,
		defaultModelName: defaultModelName,
	}, nil
}

func (mr *ModelRegistry) GetDefaultModel() (LLM, bool) {
	return mr.GetModel(mr.defaultModelName)
}

func (mr *ModelRegistry) GetModel(modelName string) (LLM, bool) {
	l, ok := mr.registry[modelName]
	return l, ok
}

// Copyright (c) 2024 Mihir Sathe
// SPDX-License-Identifier: Apache-2.0
package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/satmihir/mai/config"
)

const (
	// This seems like the only host name openai supports at the moment
	// but we will allow overriding it using the "hostname" argument in model config.
	defaultOpenAiHost = "api.openai.com"
	// Model arg for the optional hostname override.
	hostnameOverrideArg = "hostname"
	// Model arg for the API key
	apiKeyArg = "apiKey"
)

// The OpenAI implementation of the LLM interface.
type OpenAiLlm struct {
	// The model to use for the prompts against this type.
	// You have to instantiate this type once per OpenAI model you support.
	model *config.Model
	// Hostname for the OpenAI endpoint
	hostname string
	// API key to auth against the OpenAI endpoint
	apiKey string
}

type oaiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}
type oaiRequest struct {
	Model    string        `json:"model"`
	Messages []*oaiMessage `json:"messages"`
}

type oaiRespChoice struct {
	Index        uint32      `json:"index"`
	FinishReason string      `json:"finish_reason"`
	Message      *oaiMessage `json:"message"`
}
type oaiResp struct {
	Id      string           `json:"id"`
	Model   string           `json:"model"`
	Created uint64           `json:"created"`
	Choices []*oaiRespChoice `json:"choices"`
}

// Create a new instance of an OpenAI LLM implementation
func NewOpenAiLlm(model *config.Model) (*OpenAiLlm, error) {
	apiKey, ok := model.Args[apiKeyArg]
	if !ok {
		return nil, errors.Errorf("Failed to get API key from model args")
	}

	hostname := defaultOpenAiHost
	// An override is optional so no need to handle the !ok case
	override, ok := model.Args[hostnameOverrideArg]
	if ok {
		hostname = override
	}

	return &OpenAiLlm{model: model, apiKey: apiKey, hostname: hostname}, nil
}

// Complete the given prompt using our OpenAI model.
func (ol *OpenAiLlm) Complete(prompt *Prompt) (*Completion, error) {

	msgs := make([]*oaiMessage, 0)

	for _, m := range prompt.MessgeHistory {
		msgs = append(msgs, &oaiMessage{
			Role:    m.Role,
			Content: m.Content,
			Name:    m.Name,
		})
	}

	reqBody := &oaiRequest{
		Model:    ol.model.ProviderName,
		Messages: msgs,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	//fmt.Println("Sending", string(bodyBytes))

	url := fmt.Sprintf("https://%s/v1/chat/completions", ol.hostname)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ol.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		reader := resp.Body
		defer reader.Close()
		var respBodyBytes []byte
		_, _ = reader.Read(respBodyBytes)

		return nil, errors.Errorf("Failed with status [%d: %s]; Message: [%s]", resp.StatusCode, resp.Status, respBodyBytes)
	}

	reader := resp.Body
	defer reader.Close()

	var respBodyBytes []byte
	respBodyBytes, err = io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(respBodyBytes))

	var oaiResponse *oaiResp
	err = json.Unmarshal(respBodyBytes, &oaiResponse)
	if err != nil {
		return nil, err
	}

	return translateResponse(oaiResponse), nil
}

func translateResponse(resp *oaiResp) *Completion {
	choices := make([]*CompletionChoice, 0)
	for _, c := range resp.Choices {
		choices = append(choices, &CompletionChoice{
			FinishReason: c.FinishReason,
			Index:        c.Index,
			Text:         c.Message.Content,
		})
	}

	return &Completion{
		Choices: choices,
	}
}

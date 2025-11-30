package openrouter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

const Url = "https://openrouter.ai/api/v1/chat/completions"

type OpenRouterClient struct {
	token string
}

func NewOpenRouterClient(token string) *OpenRouterClient {
	return &OpenRouterClient{
		token: token,
	}
}

func (c *OpenRouterClient) getBasePath() (basePath string) {
	return Url + c.token
}

// SendPrompt send post request, accept the response and return the quote
func (c *OpenRouterClient) SendPrompt(ctx context.Context, prompt string) (string, error) {
	body := RequestBody{
		Model: "x-ai/grok-4.1-fast:free",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		ExtraBody: ThoughtsExtraBody{},
	}
	bodyJson, err := json.MarshalIndent(body, "", "\t")
	if err != nil {
		return "", fmt.Errorf("openrouter: failed marshal body %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, Url, bytes.NewReader(bodyJson))
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("openrouter: create request error %w", err)
	}
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("openrouter: failed read body %w", err)
	}
	result := gjson.GetBytes(bodyResp, `choices.0.message.content`)
	fmt.Println(result)
	defer resp.Body.Close()
	return result.Str, nil
}

package openrouter

import (
	"os"
)

type OpenRouterConfig struct {
	Token string
}

func NewOpenTouterConfig() *OpenRouterConfig {
	return &OpenRouterConfig{
		Token: os.Getenv("OPENROUTER_TOKEN"),
	}
}

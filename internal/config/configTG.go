package config

import (
	"os"
)

type TelegramConfig struct {
	ApiKey string
	ChatId string
}

func Key() *TelegramConfig {
	return &TelegramConfig{
		ApiKey: os.Getenv("BOT_TOKEN"),
		ChatId: os.Getenv("CHAT_ID"),
	}
}

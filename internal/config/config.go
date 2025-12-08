package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
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

var V *viper.Viper

func init() {
	V = viper.New()
	V.SetConfigName("config")
	V.SetConfigType("yml")
	V.AddConfigPath("./config")
	if err := V.ReadInConfig(); err != nil {
		log.Fatal("config: failed read in config %w", err)
	}

}

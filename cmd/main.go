package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"telegram_bot/internal/config"
	"telegram_bot/internal/meme"
	"telegram_bot/internal/openrouter"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

const ctxTimeout = 180 * time.Second

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл env не найден или не загружен")
	}
	botToken := os.Getenv("BOT_TOKEN")
	chatIDStr := os.Getenv("CHAT_ID")
	if botToken == "" || chatIDStr == "" {
		log.Println("Необходимые переменные окружения отсутствуют", "BOT_TOKEN", botToken, "CHAT_ID", chatIDStr)
		os.Exit(1)
	}
	cfg := openrouter.NewOpenTouterConfig()
	cfgTG := config.Key()
	m := meme.NewMeme(cfgTG, cfg)

	c := cron.New()
	_, err := c.AddFunc(config.V.GetString("time_cron"), func() {
		log.Println("Порция мемов,держим +мораль")
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()
		err := m.SendMeme(ctx)
		if err != nil {
			m.SendErr(ctx)
			panic(err)
		}
	})
	if err != nil {
		fmt.Errorf("main: failed add cron %w", err)
	}
	log.Println("Бот активен,СТАРТУЕМ!")
	c.Start()
	select {}
}

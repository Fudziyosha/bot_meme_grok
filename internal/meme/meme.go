package meme

import (
	"context"
	"fmt"
	"math/rand"
	"telegram_bot/internal/adapters"
	"telegram_bot/internal/config"
	"telegram_bot/internal/openrouter"
)

type Meme struct {
	tgClient         *adapters.TelegramClient
	chatID           string
	botToken         string
	openRouterClient *openrouter.OpenRouterClient
	openRouterConfig *openrouter.OpenRouterConfig
}

func NewMeme(cfg *config.TelegramConfig, openRouterConfig *openrouter.OpenRouterConfig) *Meme {
	return &Meme{
		tgClient:         adapters.NewTelegramClient(cfg.ApiKey),
		botToken:         cfg.ApiKey,
		chatID:           cfg.ChatId,
		openRouterConfig: openRouterConfig,
		openRouterClient: openrouter.NewOpenRouterClient(openRouterConfig.Token),
	}
}

func (m *Meme) SendErr(ctx context.Context) {
	err := m.tgClient.SendMessage(ctx, "Мемы ушли,но вы держитесь ❤️", m.chatID)
	if err != nil {
		_ = fmt.Errorf("meme: send err message tgclient failed %w", err)
	}
}

func (m *Meme) SendMeme(ctx context.Context) error {
	randQuote := m.openRouterConfig.Prompt[rand.Intn(len(m.openRouterConfig.Prompt))]
	quote, err := m.openRouterClient.SendPrompt(ctx, randQuote)
	if err != nil {
		return fmt.Errorf("meme: failed send prompt %w", err)
	}
	err = m.tgClient.SendMessage(ctx, quote, m.chatID)
	if err != nil {
		return fmt.Errorf("meme: send message tgclient failed %w", err)
	}

	return nil
}

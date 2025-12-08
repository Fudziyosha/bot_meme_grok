package adapters

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const tgApiUrl = "https://api.telegram.org"

type TelegramClient struct {
	token string
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		token: token,
	}
}

func (c *TelegramClient) getBasePath() (basePath string) {
	return tgApiUrl + "/bot" + c.token
}

func (c *TelegramClient) SendMessage(ctx context.Context, text, chatId string) error {
	tgUrl := c.getBasePath() + "/sendMessage"
	params := url.Values{}
	params.Add("chat_id", chatId)
	params.Add("text", text)
	fullURL := tgUrl + "?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("telegram_adapter: create request error %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("telegram_adapter: http request failed %w", err)
	}
	statusResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("telegram_adapter: failed read resp body status %w", err)
	}
	fmt.Println(string(statusResp))
	defer resp.Body.Close()
	return nil
}

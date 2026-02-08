package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot service wrapper.
type Bot struct {
	API *tgbotapi.BotAPI
}

// NewBot initializes a new Telegram Bot instance.
func NewBot(cfg *Config) (*Bot, error) {
	if cfg.Token == "" {
		return nil, fmt.Errorf("token is required")
	}

	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot api: %w", err)
	}

	api.Debug = cfg.Debug

	return &Bot{
		API: api,
	}, nil
}

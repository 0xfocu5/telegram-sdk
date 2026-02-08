package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// UpdateHandlerFunc defines a function to handle incoming updates.
type UpdateHandlerFunc func(update tgbotapi.Update)

// StartPolling starts polling for updates and passes them to the handler.
// This function blocks until the channel is closed.
func (b *Bot) StartPolling(timeout int, handler UpdateHandlerFunc) {
	u := tgbotapi.NewUpdate(0)
	if timeout <= 0 {
		timeout = 60
	}
	u.Timeout = timeout

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		handler(update)
	}
}

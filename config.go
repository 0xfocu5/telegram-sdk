package telegram

// Config holds the configuration for the Telegram Bot SDK.
type Config struct {
	// Token is the Telegram Bot API token.
	Token string
	// Debug enables debug logging for the bot API.
	Debug bool
	// Timeout is the timeout in seconds for long polling updates (default: 60).
	Timeout int
}

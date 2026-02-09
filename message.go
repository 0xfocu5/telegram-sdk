package telegram

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendMessage sends a text message to a chat, optionally with an inline keyboard.
func (b *Bot) SendMessage(text string, chatID int64, replyToMsgID int, disableWebPagePreview bool, keyboards ...tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	if b.API == nil {
		return tgbotapi.Message{}, fmt.Errorf("bot API is nil")
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.DisableWebPagePreview = disableWebPagePreview
	if replyToMsgID != 0 {
		msg.ReplyToMessageID = replyToMsgID
	}
	if len(keyboards) > 0 {
		msg.ReplyMarkup = keyboards[0]
	}

	return b.API.Send(msg)
}

// SendPhoto sends an image (byte slice) to a chat, with optional caption and keyboard.
func (b *Bot) SendPhoto(photoBytes []byte, chatID int64, caption string, replyToMsgID int, keyboards ...tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	if b.API == nil {
		return tgbotapi.Message{}, fmt.Errorf("bot API is nil")
	}

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture.jpg",
		Bytes: photoBytes,
	}

	photo := tgbotapi.NewPhoto(chatID, photoFileBytes)
	photo.ParseMode = tgbotapi.ModeMarkdown
	if caption != "" {
		photo.Caption = caption
	}
	if replyToMsgID != 0 {
		photo.ReplyToMessageID = replyToMsgID
	}
	if len(keyboards) > 0 {
		photo.ReplyMarkup = keyboards[0]
	}

	return b.API.Send(photo)
}

// SendPhotoFromURL sends an image from a URL, with optional caption and keyboard.
func (b *Bot) SendPhotoFromURL(photoURL string, chatID int64, caption string, replyToMsgID int, keyboards ...tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	if b.API == nil {
		return tgbotapi.Message{}, fmt.Errorf("bot API is nil")
	}

	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(photoURL))
	photo.ParseMode = tgbotapi.ModeMarkdown
	if caption != "" {
		photo.Caption = caption
	}
	if replyToMsgID != 0 {
		photo.ReplyToMessageID = replyToMsgID
	}
	if len(keyboards) > 0 {
		photo.ReplyMarkup = keyboards[0]
	}

	return b.API.Send(photo)
}

// DeleteMessage deletes a message immediately.
func (b *Bot) DeleteMessage(chatID int64, messageID int) error {
	if b.API == nil {
		return fmt.Errorf("bot API is nil")
	}

	del := tgbotapi.NewDeleteMessage(chatID, messageID)
	resp, err := b.API.Request(del)
	if err != nil {
		return err
	}
	if !resp.Ok {
		return fmt.Errorf("failed to delete message: %s", resp.Description)
	}
	return nil
}

// DeleteMessageAfter deletes a message after a specified duration (in seconds).
// This runs in a goroutine.
func (b *Bot) DeleteMessageAfter(chatID int64, messageID int, seconds int) {
	time.AfterFunc(time.Duration(seconds)*time.Second, func() {
		if err := b.DeleteMessage(chatID, messageID); err != nil {
			log.Printf("Failed to delete message %d in chat %d: %v", messageID, chatID, err)
		}
	})
}

// SendMessageAndDelete sends a message and deletes it after a delay.
func (b *Bot) SendMessageAndDelete(text string, chatID int64, seconds int) {
	msg, err := b.SendMessage(text, chatID, 0, true)
	if err != nil {
		log.Printf("Failed to send message to be deleted: %v", err)
		return
	}
	b.DeleteMessageAfter(chatID, msg.MessageID, seconds)
}

// EditMessage edits the text of a message.
func (b *Bot) EditMessage(chatID int64, messageID int, text string, disableWebPagePreview bool) error {
	if b.API == nil {
		return fmt.Errorf("bot API is nil")
	}

	edit := tgbotapi.NewEditMessageText(chatID, messageID, text)
	edit.ParseMode = tgbotapi.ModeMarkdown
	edit.DisableWebPagePreview = disableWebPagePreview
	_, err := b.API.Send(edit)
	return err
}

// EditCaption edits the caption of a photo message.
func (b *Bot) EditCaption(chatID int64, messageID int, caption string) error {
	if b.API == nil {
		return fmt.Errorf("bot API is nil")
	}

	edit := tgbotapi.NewEditMessageCaption(chatID, messageID, caption)
	edit.ParseMode = tgbotapi.ModeMarkdown
	_, err := b.API.Send(edit)
	return err
}

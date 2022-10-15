package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Errors.Default)
	switch err {
	case errInvalidURL:
		msg.Text = b.messages.Errors.InvalidURL
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = b.messages.Errors.Unauthorized
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = b.messages.Errors.UnableToSeve
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}

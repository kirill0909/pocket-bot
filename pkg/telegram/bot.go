package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kirill0909/pocket-bot/pkg/config"
	"github.com/kirill0909/pocket-bot/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
)

// Bot just wrapper of BotAPI
type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	redirectURL     string
	tokenRepository repository.TokenRepository

	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirectURL string,
	tokenRepository repository.TokenRepository, messages config.Messages) *Bot {
	return &Bot{
		bot:             bot,
		pocketClient:    pocketClient,
		redirectURL:     redirectURL,
		tokenRepository: tokenRepository,
		messages:        messages,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					b.handleError(update.Message.Chat.ID, err)
				}
				continue
			}

			if err := b.handleMessage(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

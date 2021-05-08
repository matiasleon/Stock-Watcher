package bot

import (
	"log"

	"api/main.go/api/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot represents an bot entity that will listen incoming messages
type Bot struct {
	botClient    *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
}

// CreateBot creates an instance of Bot
func Create() (*Bot, error) {
	botClient, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		return nil, err
	}
	newBot := &Bot{botClient: botClient}
	newBot.setup()

	return newBot, nil
}

// setup initialize necessary properties for work
func (b *Bot) setup() {
	b.botClient.Debug = true
	log.Printf("Authorized on account %s", b.botClient.Self.UserName)

	// create update config
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	b.updateConfig = updateConfig
}

// Listen listens incoming messages
func (b *Bot) Listen() {
	updates, err := b.botClient.GetUpdatesChan(b.updateConfig)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hola amor")
		msg.ReplyToMessageID = update.Message.MessageID

		b.botClient.Send(msg)
	}
}

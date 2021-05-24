package bot

import (
	"log"
	"stocktracker/api/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type interpreter interface {
	Interpret(name, messageText string) (string, error)
}

// Bot represents an bot entity that will listen incoming messages
type Bot struct {
	botClient    *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	interpreter  interpreter
}

// CreateBot creates an instance of Bot
func Create(interpreter interpreter) (*Bot, error) {
	botClient, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		return nil, err
	}
	newBot := &Bot{botClient: botClient, interpreter: interpreter}
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

		name := update.Message.From.FirstName + " " + update.Message.From.LastName
		message, err := b.interpreter.Interpret(name, update.Message.Text)
		if err != nil {
			log.Panic(err)
		}

		b.sendMessage(update.Message, message)
	}
}

func (b *Bot) sendMessage(incomingMessage *tgbotapi.Message, message string) {
	msg := tgbotapi.NewMessage(incomingMessage.Chat.ID, message)
	msg.ReplyToMessageID = incomingMessage.MessageID

	b.botClient.Send(msg)
}

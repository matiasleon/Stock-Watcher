package bot

type interpreter interface {
	Interpret(name, messageText string) (string, error)
}

// Bot represents an bot entity that will listen incoming messages
type Bot struct {
	interpreter interpreter
}

// CreateBot creates an instance of Bot
func Create(interpreter interpreter) (*Bot, error) {
	newBot := &Bot{}
	newBot.setup()

	return newBot, nil
}

// setup initialize necessary properties for work
func (b *Bot) setup() {
}

// Listen listens incoming messages
func (b *Bot) Listen() {

}

func (b *Bot) sendMessage() {
}

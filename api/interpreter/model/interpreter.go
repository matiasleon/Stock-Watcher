package interpretermodel

import "fmt"

// interpreter holds the logic to interprete messages
type interpreter struct {
}

func New() *interpreter {
	return &interpreter{}
}

// processIncomingMessage interpret incoming messages
func (i *interpreter) Interpret(name, messageText string) (string, error) {
	switch messageText {
	case "/start":
		return fmt.Sprintf("Hola %s, bienvenido al bot de alertas! Te voy a estar ayudando", name), nil

	case "/hello":
		return fmt.Sprintf("Hola %s, como estas? Estoy preparado para ayudarte :)", name), nil
	default:
		return fmt.Sprint("Lo siento, no entiendo lo que me dices"), nil
	}
}

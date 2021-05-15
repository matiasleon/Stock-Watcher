package main

import (
	"log"

	interpretermodel "api/main.go/api/interpreter/model"

	"api/main.go/api/bot"
)

func main() {

	interpreter := interpretermodel.New()

	// create bot
	stockBot, err := bot.Create(interpreter)
	if err != nil {
		log.Panic(err)
	}

	stockBot.Listen()
}

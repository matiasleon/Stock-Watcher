package main

import (
	"log"
	"stocktracker/api/bot"
	interpretermodel "stocktracker/api/interpreter/model"
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

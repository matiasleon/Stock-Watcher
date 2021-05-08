package main

import (
	"log"

	"api/main.go/api/bot"
)

func main() {

	// create bot
	stockBot, err := bot.Create()
	if err != nil {
		log.Panic(err)
	}

	stockBot.Listen()
}

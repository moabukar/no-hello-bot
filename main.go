package main

import (
	"fmt"

	"github.com/moabukar/discord-bot-go/bot"
	"github.com/moabukar/discord-bot-go/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error)
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}

package main

import (
	"fmt"

	"github.com/moabukar/discord-bot-go/config"
	"github.com/moabukar/discord-bot-go/pkg/bot"
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

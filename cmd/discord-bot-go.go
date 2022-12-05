package main

import (
	"fmt"

	"github.com/moabukar/no-hello-bot/config"
	"github.com/moabukar/no-hello-bot/pkg/bot"
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

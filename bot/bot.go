package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/moabukar/discord-bot-go/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {

	goBot, err := discordgo.New("Bot" + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")

}

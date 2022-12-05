package bot

import (
	"fmt"

	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/moabukar/no-hello-bot/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

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

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// if m.Author.ID == BotID {
	// 	return
	// }

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pongg!")
	}

	if m.Content == "salaam!" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "walaukum salaam!")
	}

	if m.Content == "assalamu alaykum" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "walaukum salaam wrwb!")
	}

	// If the message is "hello"
	if strings.ToLower(m.Content) == "hello" {
		// Send a message with nohello.com
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "hi" {
		// Send a message with nohello.com
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.Contains(m.Content, s.State.User.Mention()) {
		// Send a message with nohello.com when tagged.
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "nh" {
		// Send a message with nohello.com when "nh" is used
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "salaam" {
		// Send a prompt when a salaam is written
		s.ChannelMessageSend(m.ChannelID, "Walaykum salaam! How can we help 🙂 If you've got a query, make sure to put in as much detail as possible!")
	}

	if strings.ToLower(m.Content) == "salam" {
		// Salaam variation 2.0
		s.ChannelMessageSend(m.ChannelID, "Walaykum salaam! How can we help 🙂 If you've got a query, make sure to put in as much detail as possible!")
	}

	if strings.ToLower(m.Content) == "assalamu alaykum" {
		// Salaam variation 3.0
		s.ChannelMessageSend(m.ChannelID, "walaukum salaam wrwb!")
	}

	if strings.ToLower(m.Content) == "nc" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, "Please give more context to your question!! Outline the issue at hand; giving enough detail and background for someone to help")
	}
}

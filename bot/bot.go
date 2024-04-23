package bot

import (
	"fmt"

	"github.com/SteakBarbare/GoBot-Example/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
	fmt.Println(err.Error())
	return
	}

	u, err := goBot.User("@me")

	if err != nil {
	fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
	fmt.Println(err.Error())
	return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
	return
	}

	// if m.content contains botid (Mentions) and "ping" then send "pong!"
	if m.Content == "cc" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "sv ?")
	}

	if m.Content == config.BotPrefix+"Sauce" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Executing command Sauce")
	}

	if m.Author.ID == "142335378064408585"{
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ta gueule")
	}
	if m.Author.ID == "303133529485869058"{
		_, _ = s.ChannelMessageSend(m.ChannelID, "Tu es hystérique, vas te calmer")
	}
}
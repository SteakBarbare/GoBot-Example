package bot

import (
	"fmt"

	"github.com/SteakBarbare/GoBot-Example/config"
	"github.com/bwmarrin/discordgo"
)

var botId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot" + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botUser, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botId = botUser.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botId {
		return
	}

	if m.Content == "cc" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "sv ?")
	}

	if m.Content == config.BotPrefix+"Sauce" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Executing command Sauce")
	}
}
package config

import (
	"os"
)

var (
	Token     string
	BotPrefix string
)

func LoadConfig() error {
	
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	return nil
}
package main

import (
	"fmt"

	"github.com/SteakBarbare/GoBot-Example/bot"
	"github.com/SteakBarbare/GoBot-Example/config"
)

func main() {
	err := config.LoadConfig()

	if(err != nil){
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
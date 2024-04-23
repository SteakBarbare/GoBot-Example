package main

import (
	"fmt"

	"github.com/SteakBarbare/GoBot-Example/bot"
	"github.com/SteakBarbare/GoBot-Example/config"
)

func main() {
	fmt.Println("Bipute")
	err := config.LoadConfig()

	if(err != nil){
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Tripute")
	bot.Start()

	<-make(chan struct{})

	return
}
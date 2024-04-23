package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	token     string
	botPrefix string

	config *Config
)

type Config struct {
	token     string `json:"TOKEN"`
	botPrefix string `json:"BOT_PREFIX"`
}

func LoadConfig() error {
	fmt.Println("Loading config file")
	file, err := os.ReadFile("./config.json")
	
	if(err != nil){
		fmt.Println("Error while loading the file")
		return err
	}

	fmt.Println("Converting file to config struct")
	err = json.Unmarshal(file, &config)

	if(err != nil){
		fmt.Println("Error while converting the file")
		return err
	}

	token = config.token
	botPrefix = config.botPrefix

	return nil
}
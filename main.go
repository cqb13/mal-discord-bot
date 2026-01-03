package main

import (
	"dev/cqb13/mal-bot/bot"
	"dev/cqb13/mal-bot/bot/commands/addonList"
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env file: ", err)
		os.Exit(1)
	}

	var key string = os.Getenv("KEY")
	var githubToken string = os.Getenv("GITHUB_TOKEN")

	utils.InitDefaultHeaders(githubToken)

	_, err = addonList.UseList()
	if err != nil {
		fmt.Printf("Failed to load initial list: %s\n", err)
		os.Exit(1)
	}

	bot.BotToken = key
	bot.Run()
}

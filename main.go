package main

import (
	"dev/cqb13/mal-bot/bot"
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

	var key string = os.Getenv("TEST_BOT")
	var githubToken string = os.Getenv("GITHUB_TOKEN")

	utils.InitDefaultHeaders(githubToken)

	bot.BotToken = key
	bot.Run()
}

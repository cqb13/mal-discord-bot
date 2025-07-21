package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"dev/cqb13/mal-bot/bot/commands"
	"dev/cqb13/mal-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	session, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	session.AddHandler(commands.HandleInteractions)
	utils.Log("say command added")
	session.AddHandler(commands.HandleSayCommand)

	session.Open()
	defer session.Close()

	appID := session.State.User.ID
	guildID := ""

	for _, cmd := range commands.Commands {
		_, err := session.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			utils.Log(fmt.Sprintf("Failed to create command %s: %v", cmd.Name, err))
		}
		utils.Log(fmt.Sprintf("%s command added", cmd.Name))
	}

	utils.Log("Bot Running")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

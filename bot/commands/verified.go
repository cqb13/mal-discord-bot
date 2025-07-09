package commands

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var VerifiedCommand = &discordgo.ApplicationCommand{
	Name:        "verified",
	Description: "Lists all current verified addons, updates hourly.",
}

var verifiedAddons string
var fetchTime time.Time

func handleVerified(s *discordgo.Session, i *discordgo.InteractionCreate) {
	currentTime := time.Now()

	if currentTime.Sub(fetchTime) > time.Hour {
		fetched, err := getVerifiedAddons()
		if err != nil {
			utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
			return
		}
		verifiedAddons = fetched
		fetchTime = currentTime
	}

	embed := &discordgo.MessageEmbed{
		Title: "Verified Addons",
		Color: utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "",
				Value: verifiedAddons,
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

func getVerifiedAddons() (string, error) {
	url := "https://raw.githubusercontent.com/cqb13/meteor-addon-scanner/main/verified.txt"
	bytes, err := utils.MakeGetRequest(url)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

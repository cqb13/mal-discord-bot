package commands

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotCommand = &discordgo.ApplicationCommand{
	Name:        "bot",
	Description: "Shows information about the bot",
}

func handleBot(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:     "Meteor Addon List Bot",
		URL:       "https://github.com/cqb13/mal-discord-bot",
		Color:     utils.EmbedColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://www.meteoraddons.com/icon.png"},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Start Time",
				Value: utils.TimeToPrettyStr(utils.StartTime),
			},
			{
				Name:  "Up Time",
				Value: utils.CalculateUptime(),
			},
			{
				Name:  "Commands",
				Value: fmt.Sprintf("%d", len(Commands)),
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

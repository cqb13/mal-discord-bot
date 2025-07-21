package commands

import (
	"dev/cqb13/mal-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var DeveloperInfoCommand = &discordgo.ApplicationCommand{
	Name:        "developer-info",
	Description: "Provides information for addon developers",
}

func handleDeveloperInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Developer Info",
		Description: "Information for addon developers.",
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Custom Properties",
				Value: "The scanner automatically pulls info from GitHub, but it might not always be accurate or exactly how you want it. To fix or customize that data, you can manually add your own values.\n\nTo do that, create the file `meteor-addon-list.json` in the root directory of your addon, and add the fields you wish to overide:\n```json\n{\n\t\"description\": \"A short description of your addon.\",\n\t\"supported_versions\": [\"1.21.7\", \"1.21.8\"],\n\t\"icon\": \"https://example.com/icon.png\",\n\t\"discord\": \"https://discord.gg/yourserver\",\n\t\"homepage\": \"https://example.com\"\n}\n```",
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, true, "")
}

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
				Value: "The scanner automatically pulls info from GitHub, but it might not always be accurate or exactly how you want it. To fix or customize that data, you can manually add your own values.\n\nTo do that, create the file `meteor-addon-list.json` in the root directory of your addon, and add the fields you wish to overide:\n```json\n{\n\t\"description\": \"A short description of your addon.\",\n\t\"tags\": [\n\t\t\"PvP\",\n\t\t\"Utility\",\n\t\t\"Theme\",\n\t\t\"...\"\n\t],\n\t\"supported_versions\": [\n\t\t\"1.21.7\",\n\t\t\"1.21.8\"\n\t],\n\t\"icon\": \"https://meteoraddons.com/icon.png\",\n\t\"discord\": \"https://discord.gg/XU7Y9G46KD\",\n\t\"homepage\": \"https://www.meteoraddons.com\"\n}\n```",
			},
			{
				Name:  "Supported Tags",
				Value: "- Pvp\n- Utility\n- Theme\n- Render\n- Movement\n- Building\n- World\n- Misc\n- QoL\n- Exploit\n- Fun\n- Automation\n",
			},
		},
	}

	err := utils.InteractionRespondEmbed(embed, s, i.Interaction, true, "")
	if err != nil {
		utils.LogCmdResponseFailiure(i, err.Error())
	}
}

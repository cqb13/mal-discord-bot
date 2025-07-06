package commands

import (
	"dev/cqb13/mal-bot/utils"

	"fmt"
	"github.com/bwmarrin/discordgo"
)

var NotifyCommand = &discordgo.ApplicationCommand{
	Name:        "notify",
	Description: "Notifies when a new addon is verified.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "name",
			Description: "Name of the addon",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "author",
			Description: "Author of the addon",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "repo",
			Description: "GitHub link to the addon repository: user/repo",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "description",
			Description: "Short description of what your addon does",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
		},
	},
}

func handleNotify(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !utils.IsOwner(i) {
		utils.InteractionRespondText("You do not have permission to run this command.", s, i.Interaction, true, "")
		return
	}

	options := i.ApplicationCommandData().Options

	var name, author, repo, description string

	for _, opt := range options {
		switch opt.Name {
		case "name":
			name = opt.StringValue()
		case "author":
			author = opt.StringValue()
		case "repo":
			repo = opt.StringValue()
		case "description":
			description = opt.StringValue()
		}
	}

	descriptionLine := ""

	if description != "" {
		descriptionLine = fmt.Sprintf("**Description:** %s", description)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "New Addon Verified!",
		Description: fmt.Sprintf("<@&%s> A new addon has been verified!\n**Name:** %s\n**Author:** %s\n**Repo:** https://github.com/%s\n%s", utils.NotifiedRoleId, name, author, repo, descriptionLine),
		Color:       utils.EmbedColor,
	}

	utils.SendToChannelEmbed(embed, s, "1377120291452358690", utils.NotifiedRoleId)
	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, utils.NotifiedRoleId)
}

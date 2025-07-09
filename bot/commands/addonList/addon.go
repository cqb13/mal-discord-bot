package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var AddonCommand = &discordgo.ApplicationCommand{
	Name:        "addon",
	Description: "Gives information about an addon in the list, updates hourly.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "owner",
			Description: "Owner of the repository",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "repo",
			Description: "Name of the repository",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func HandleAddon(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	var owner, repo string

	for _, opt := range options {
		switch opt.Name {
		case "owner":
			owner = opt.StringValue()
		case "repo":
			repo = opt.StringValue()
		}
	}

	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	fullName := strings.ToLower(fmt.Sprintf("%s/%s", owner, repo))

	var addon *Addon = nil

	for _, a := range list {
		if fullName == strings.ToLower(a.Repo.Id) {
			addon = &a
			break
		}
	}

	if addon == nil {
		utils.InteractionRespondText("This addon is not in the list.", s, i.Interaction, true, "")
		return
	}

	createdAt, err := utils.RFC3339StrToPrettyStr(addon.Repo.CreationDate)
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	lastUpdate, err := utils.RFC3339StrToPrettyStr(addon.Repo.LastUpdate)
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       addon.Name,
		Description: addon.Description,
		Color:       utils.EmbedColor,
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: addon.Links.Icon},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "MC Version",
				Value:  addon.McVersion,
				Inline: true,
			},
			{
				Name:   "Verified",
				Value:  utils.Ternary(addon.Verified, "Yes", "No"),
				Inline: true,
			},
			{
				Name:   "Feature Count",
				Value:  strconv.Itoa(addon.FeatureCount),
				Inline: true,
			},
			{
				Name:   "Stars",
				Value:  strconv.Itoa(addon.Repo.Stars),
				Inline: true,
			},
			{
				Name:   "Downloads",
				Value:  strconv.Itoa(addon.Repo.Downloads),
				Inline: true,
			},
			{
				Name:   "Fork",
				Value:  utils.Ternary(addon.Repo.Fork, "True", "False"),
				Inline: true,
			},
			{
				Name:   "Archived",
				Value:  utils.Ternary(addon.Repo.Fork, "True", "False"),
				Inline: true,
			},
			{
				Name:   "Created On",
				Value:  createdAt,
				Inline: false,
			},
			{
				Name:   "Latest Push",
				Value:  lastUpdate,
				Inline: false,
			},
			{
				Name:   "Authors",
				Value:  strings.Join(addon.Authors, ", "),
				Inline: false,
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

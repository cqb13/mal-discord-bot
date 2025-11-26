package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var FindCommand = &discordgo.ApplicationCommand{
	Name:        "find",
	Description: "Find a feature in addons",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "name",
			Description: "Name of the feature",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "type",
			Description: "The type of feature",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "All",
					Value: "all",
				},
				{
					Name:  "Module",
					Value: "module",
				},
				{
					Name:  "Command",
					Value: "command",
				},
				{
					Name:  "HUD",
					Value: "hud",
				},
			},
		},
	},
}

func HandleFind(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var name string
	var featureType string = "all"

	for _, opt := range i.ApplicationCommandData().Options {
		switch opt.Name {
		case "name":
			name = strings.ToLower(opt.StringValue())
		case "type":
			featureType = opt.StringValue()
		}
	}

	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	var haveFeature []Addon

	for _, addon := range list {
		if (featureType == "all" || featureType == "module") && slices.ContainsFunc(addon.Features.Modules, func(feature string) bool {
			return strings.ToLower(feature) == name
		}) {
			haveFeature = append(haveFeature, addon)

			continue
		}

		if (featureType == "all" || featureType == "command") && slices.ContainsFunc(addon.Features.Commands, func(feature string) bool {
			return strings.ToLower(feature) == name
		}) {
			haveFeature = append(haveFeature, addon)

			continue
		}

		if (featureType == "all" || featureType == "hud") && slices.ContainsFunc(addon.Features.HudElements, func(feature string) bool {
			return strings.ToLower(feature) == name
		}) {
			haveFeature = append(haveFeature, addon)

			continue
		}
	}

	sort.Slice(haveFeature, func(i, j int) bool {
		return haveFeature[i].Repo.Stars > haveFeature[j].Repo.Stars
	})

	withFeature := ""

	for _, addon := range haveFeature {
		withFeature += fmt.Sprintf("- **[%s](%s)** - %s\n", addon.Name, addon.Links.Github, addon.McVersion)
	}

	if len(haveFeature) == 0 {
		withFeature = "None"
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Addons with %s", name),
		Description: fmt.Sprintf("Meteor addon list as of %s", utils.TimeToPrettyStr(FetchTime)),
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "",
				Value: withFeature,
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

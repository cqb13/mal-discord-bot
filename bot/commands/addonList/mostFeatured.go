package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var MostFeaturedCommand = &discordgo.ApplicationCommand{
	Name:        "most-featured",
	Description: "Gives the 10 addons with the most features from the addon list, updates hourly.",
}

func HandleMostFeatured(s *discordgo.Session, i *discordgo.InteractionCreate) {
	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		utils.LogCmdError(i, err.Error())
		return
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Repo.Fork {
			return false
		}

		return list[i].Features.FeatureCount > list[j].Features.FeatureCount
	})

	var top10Str strings.Builder

	for i := range 10 {
		fmt.Fprintf(&top10Str, "- **%s** - %d features\n", list[i].Name, list[i].Features.FeatureCount)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Most Full-Featured Addons",
		Description: fmt.Sprintf("Meteor addon list as of %s", utils.TimeToPrettyStr(FetchTime)),
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "",
				Value: top10Str.String(),
			},
		},
	}

	err = utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
	if err != nil {
		utils.LogCmdResponseFailiure(i, err.Error())
	}
}

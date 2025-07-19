package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"sort"

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
		return
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Repo.Fork {
			return false
		}

		return list[i].Features.FeatureCount > list[j].Features.FeatureCount
	})

	top10Str := ""

	for i := range 10 {
		top10Str += fmt.Sprintf("- **%s** - %d features\n", list[i].Name, list[i].Features.FeatureCount)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Most Full-Featured Addons",
		Description: fmt.Sprintf("Meteor addon list as of %s", utils.TimeToPrettyStr(FetchTime)),
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "",
				Value: top10Str,
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

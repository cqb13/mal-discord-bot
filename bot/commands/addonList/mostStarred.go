package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var MostStarredCommand = &discordgo.ApplicationCommand{
	Name:        "most-starred",
	Description: "Gives the 10 addons with the most stars from the addon list, updates hourly.",
}

func HandleMostStarred(s *discordgo.Session, i *discordgo.InteractionCreate) {
	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Repo.Stars > list[j].Repo.Stars
	})

	var top10Str strings.Builder

	for i := range 10 {
		fmt.Fprintf(&top10Str, "- **%s** - %d stars\n", list[i].Name, list[i].Repo.Stars)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Most Starred Addons",
		Description: fmt.Sprintf("Meteor addon list as of %s", utils.TimeToPrettyStr(FetchTime)),
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "",
				Value: top10Str.String(),
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

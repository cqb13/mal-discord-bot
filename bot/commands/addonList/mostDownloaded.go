package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var MostDownloadedCommand = &discordgo.ApplicationCommand{
	Name:        "most-downloaded",
	Description: "Gives the 10 addons with the most downloads from the addon list, updates hourly.",
}

func HandleMostDownloaded(s *discordgo.Session, i *discordgo.InteractionCreate) {
	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		utils.LogCmdError(i, err.Error())
		return
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Repo.Downloads > list[j].Repo.Downloads
	})

	var top10Str strings.Builder

	for i := range 10 {
		fmt.Fprintf(&top10Str, "- **%s** - %d downloads\n", list[i].Name, list[i].Repo.Downloads)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Most Downloaded Addons",
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

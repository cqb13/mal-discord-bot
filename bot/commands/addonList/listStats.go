package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var ListStatsCommand = &discordgo.ApplicationCommand{
	Name:        "list-stats",
	Description: "Gives details about the addon list, updates hourly.",
}

func HandleListStats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	totalAddons := len(list)
	verifiedAddons := 0
	forkedAddons := 0
	archivedAddons := 0
	var mostStarredAddon Addon
	var mostDownloadedAddon Addon

	for _, addon := range list {
		if addon.Verified {
			verifiedAddons++
		}

		if addon.Repo.Fork {
			forkedAddons++
		}

		if addon.Repo.Archived {
			archivedAddons++
		}

		if addon.Repo.Downloads > mostDownloadedAddon.Repo.Downloads {
			mostDownloadedAddon = addon
		}

		if addon.Repo.Stars > mostStarredAddon.Repo.Stars {
			mostStarredAddon = addon
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       "List Stats",
		Description: fmt.Sprintf("Meteor addon list as of %s", utils.TimeToPrettyStr(FetchTime)),
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Total Addons",
				Value:  strconv.Itoa(totalAddons),
				Inline: true,
			},
			{
				Name:   "Verified Addons",
				Value:  strconv.Itoa(verifiedAddons),
				Inline: true,
			},
			{
				Name:   "Forked Addons",
				Value:  strconv.Itoa(forkedAddons),
				Inline: true,
			},
			{
				Name:   "Archived Addons",
				Value:  strconv.Itoa(archivedAddons),
				Inline: true,
			},
			{
				Name:  "Most Starred Addon",
				Value: fmt.Sprintf("%s by %s", mostStarredAddon.Name, mostStarredAddon.Repo.Owner),
			},
			{
				Name:  "Most Downloaded Addon",
				Value: fmt.Sprintf("%s by %s", mostDownloadedAddon.Name, mostDownloadedAddon.Repo.Owner),
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

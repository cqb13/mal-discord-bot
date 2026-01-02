package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var VerifiedCommand = &discordgo.ApplicationCommand{
	Name:        "verified",
	Description: "Lists all current verified addons, updates hourly.",
}

func HandleVerified(s *discordgo.Session, i *discordgo.InteractionCreate) {
	list, err := UseList()
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	var value strings.Builder

	for _, addon := range list {
		if addon.Verified {
			fmt.Fprintf(&value, "- **%s** by %s\n", addon.Name, addon.Repo.Owner)
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Verified Addons",
		Description: fmt.Sprintf("Meteor addon list as of %s", utils.TimeToPrettyStr(FetchTime)),
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "",
				Value: value.String(),
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

package commands

import (
	"slices"

	"dev/cqb13/mal-bot/utils"
	"github.com/bwmarrin/discordgo"
)

var NotifiedCommand = &discordgo.ApplicationCommand{
	Name:        "notified",
	Description: "Gives you the notified role, meaning you will be notified when addons are verified",
}

func handleNotified(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if slices.Contains(i.Member.Roles, utils.NotifiedRoleId) {
		s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, utils.NotifiedRoleId)
		utils.InteractionRespondText("You will no longer be notified when new addons are verified.", s, i.Interaction, true, "")
	} else {
		s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, utils.NotifiedRoleId)
		utils.InteractionRespondText("You will now be notified when new addons are verified.", s, i.Interaction, true, "")
	}
}

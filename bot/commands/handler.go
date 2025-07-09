package commands

import (
	"dev/cqb13/mal-bot/bot/commands/addonList"
	"dev/cqb13/mal-bot/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	AboutCommand,
	addonList.AddonCommand,
	addonList.ListStatsCommand,
	addonList.MostDownloadedCommand,
	addonList.MostStarredCommand,
	NotifiedCommand,
	NotifyCommand,
	RepoCommand,
	addonList.VerifiedCommand,
}

func HandleInteractions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	utils.Log(fmt.Sprintf("%s ran the %s command", i.Member.User.Username, i.ApplicationCommandData().Name))
	switch i.ApplicationCommandData().Name {
	case "about":
		handleAbout(s, i)
		return
	case "addon":
		addonList.HandleAddon(s, i)
		return
	case "list-stats":
		addonList.HandleListStats(s, i)
		return
	case "most-downloaded":
		addonList.HandleMostDownloaded(s, i)
		return
	case "most-starred":
		addonList.HandleMostStarred(s, i)
		return
	case "notified":
		handleNotified(s, i)
		return
	case "notify":
		handleNotify(s, i)
		return
	case "repo":
		handleRepo(s, i)
		return
	case "verified":
		addonList.HandleVerified(s, i)
		return
	}
}

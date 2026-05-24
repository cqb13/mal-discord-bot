package commands

import (
	"dev/cqb13/mal-bot/bot/commands/addonList"
	"dev/cqb13/mal-bot/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	AboutCommand,
	BotCommand,
	addonList.AddonCommand,
	addonList.ListStatsCommand,
	addonList.MostDownloadedCommand,
	addonList.MostFeaturedCommand,
	addonList.MostStarredCommand,
	NotifiedCommand,
	NotifyCommand,
	RepoCommand,
	addonList.VerifiedCommand,
	RescanCommand,
}

func HandleInteractions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	utils.Log(fmt.Sprintf("%s ran the %s command", i.Member.User.Username, i.ApplicationCommandData().Name))
	switch i.ApplicationCommandData().Name {
	case "about":
		handleAbout(s, i)
	case "addon":
		addonList.HandleAddon(s, i)
	case "bot":
		handleBot(s, i)
	case "list-stats":
		addonList.HandleListStats(s, i)
	case "most-downloaded":
		addonList.HandleMostDownloaded(s, i)
	case "most-featured":
		addonList.HandleMostFeatured(s, i)
	case "most-starred":
		addonList.HandleMostStarred(s, i)
	case "notified":
		handleNotified(s, i)
	case "notify":
		handleNotify(s, i)
	case "repo":
		handleRepo(s, i)
	case "verified":
		addonList.HandleVerified(s, i)
	case "rescan":
		handleRescan(s, i)
	}
}

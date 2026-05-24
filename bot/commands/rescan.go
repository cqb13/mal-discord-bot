package commands

import (
	"encoding/json"
	"fmt"
	"slices"
	"sync"
	"time"

	"dev/cqb13/mal-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	lastScanTime = make(map[string]time.Time)
	scanMu       sync.Mutex
)

var RescanCommand = &discordgo.ApplicationCommand{
	Name:        "rescan",
	Description: "Allows addon developers to trigger a list rescan once per day",
}

func handleRescan(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.GuildID != utils.MainServerId {
		err := utils.InteractionRespondText("This command can not be run outside of the Meteor Addon List server: https://discord.gg/XU7Y9G46KD", s, i.Interaction, true, "")
		if err != nil {
			utils.LogCmdResponseFailiure(i, err.Error())
		}
		return
	}

	if !slices.Contains(i.Member.Roles, utils.VerifiedDeveloperRoleID) {
		err := utils.InteractionRespondText("You must be a verified addon developer to use this command", s, i.Interaction, true, "")
		if err != nil {
			utils.LogCmdResponseFailiure(i, err.Error())
		}
		return
	}

	userID := i.Member.User.ID

	scanMu.Lock()
	last, exists := lastScanTime[userID]
	now := time.Now().UTC()
	todayStart := now.Truncate(24 * time.Hour)
	if exists && last.After(todayStart) {
		scanMu.Unlock()
		nextReset := todayStart.Add(24 * time.Hour)
		err := utils.InteractionRespondText(
			fmt.Sprintf("You've already used your daily rescan. Next scan available <t:%d:R>.", nextReset.Unix()),
			s, i.Interaction, true, "",
		)
		if err != nil {
			utils.LogCmdResponseFailiure(i, err.Error())
		}
		return
	}
	lastScanTime[userID] = now
	scanMu.Unlock()

	url := "https://api.github.com/repos/cqb13/meteor-addon-scanner/actions/workflows/scan.yml/dispatches"

	payload, _ := json.Marshal(map[string]string{"ref": "main"})

	_, err := utils.MakePostRequest(url, payload)
	if err != nil {
		utils.InteractionRespondText("Failed to trigger rescan. Please try again later.", s, i.Interaction, true, "")
		utils.LogCmdResponseFailiure(i, err.Error())
		return
	}

	err = utils.InteractionRespondText("Rescan triggered successfully! The addon list will be updated shortly.", s, i.Interaction, false, "")
	if err != nil {
		utils.LogCmdResponseFailiure(i, err.Error())
	}
}

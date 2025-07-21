package commands

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HandleSayCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!say ") {
		return
	}
	utils.Log(fmt.Sprintf("%s ran the say command", m.Author.Username))

	if m.Author.ID != utils.OwnerID {
		return
	}

	sayContent := strings.TrimSpace(m.Content[len("!say "):])

	_ = s.ChannelMessageDelete(m.ChannelID, m.ID)

	_, _ = s.ChannelMessageSend(m.ChannelID, sayContent)
}

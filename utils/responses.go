package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const EmbedColor = 0xdab2ff

func InteractionRespondText(response string, session *discordgo.Session, interaction *discordgo.Interaction, ephemeral bool, pingRoleID string) error {
	flags := discordgo.MessageFlags(0)
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	allowedMentions := &discordgo.MessageAllowedMentions{}
	if pingRoleID != "" {
		allowedMentions.Roles = []string{pingRoleID}
	}

	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         response,
			AllowedMentions: allowedMentions,
			Flags:           flags,
		},
	})
}

func InteractionRespondEmbed(embed *discordgo.MessageEmbed, session *discordgo.Session, interaction *discordgo.Interaction, ephemeral bool, pingRoleID string) error {
	flags := discordgo.MessageFlags(0)
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	content := ""
	allowedMentions := &discordgo.MessageAllowedMentions{}

	if pingRoleID != "" {
		content = fmt.Sprintf("<@&%s>\u200B", pingRoleID)
		allowedMentions.Roles = []string{pingRoleID}
	}

	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         content,
			Embeds:          []*discordgo.MessageEmbed{embed},
			AllowedMentions: allowedMentions,
			Flags:           flags,
		},
	})
}

func SendToChannelText(message string, session *discordgo.Session, channelID string, pingRoleID string) error {
	allowedMentions := &discordgo.MessageAllowedMentions{}
	if pingRoleID != "" {
		allowedMentions.Roles = []string{pingRoleID}
	}

	_, err := session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content:         message,
		AllowedMentions: allowedMentions,
	})

	return err
}

func SendToChannelEmbed(embed *discordgo.MessageEmbed, session *discordgo.Session, channelID string, pingRoleID string) error {
	content := ""
	allowedMentions := &discordgo.MessageAllowedMentions{}

	if pingRoleID != "" {
		content = fmt.Sprintf("<@&%s>\u200B", pingRoleID)
		allowedMentions.Roles = []string{pingRoleID}
	}

	_, err := session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content:         content,
		Embeds:          []*discordgo.MessageEmbed{embed},
		AllowedMentions: allowedMentions,
	})

	return err
}

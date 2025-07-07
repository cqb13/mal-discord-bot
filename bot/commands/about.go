package commands

import (
	"dev/cqb13/mal-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var AboutCommand = &discordgo.ApplicationCommand{
	Name:        "about",
	Description: "Sends the about blurb",
}

func handleAbout(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "**Meteor Addon List**",
		Description: "An ever updating list of free and open-source Meteor Client addons.",
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "",
				Value:  "**[Browse Addons](https://www.meteoraddons.com)**",
				Inline: false,
			},
			{
				Name:  "How to Get Your Addon Verified",
				Value: "Create a post in **verification-requests** with:\n- GitHub Repository Link\n- Short description of what your addon does.",
			},
			{
				Name:  "What Will Prevent Verification",
				Value: "- Harmful features (e.g. backdoors, coordinate leaks)\n- Obfuscated/unreadable code\n- Broken or non-functional addons",
			},
			{
				Name:  "",
				Value: "*It may take some time to verify your addon. You will be notified when a decision is made.*",
			},
			{
				Name:  "Contributing",
				Value: "Feel free help improve Meteor Addon List by contributing to the [website](https://github.com/cqb13/meteor-addons) and [scanner](https://github.com/cqb13/meteor-addon-scanner).",
			},
			{
				Name:  "Rules",
				Value: "1. Follow the Discord [Terms of Service](https://discord.com/terms) and [Community Guidelines](https://discord.com/guidelines).\n2. Do not attempt to spam.\n3. Do not advertise outside of the appropriate channel.",
			},
			{
				Name:  "",
				Value: "*Run **/notified** to stay updated on all new verified addons*",
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, !utils.IsOwner(i), "")
}

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
		Title:       "Meteor Addon List",
		Description: "An ever updating list of free and open-source Meteor Client addons.",
		Color:       utils.EmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "",
				Value:  "**[Browse Addons](https://meteoraddons.com)**",
				Inline: false,
			},
			{
				Name:  "Rules",
				Value: "1. Follow the Discord [Terms of Service](https://discord.com/terms) and [Community Guidelines](https://discord.com/guidelines).\n2. Do not attempt to spam.\n3. Do not advertise outside of the appropriate channel.",
			},
			{
				Name:  "Developer Info",
				Value: "If you are an addon developer, refer to the [about page](http://meteoraddons.com/about#developer-info) on the website to see how you can get the most out of the Meteor Addon List.",
			},
			{
				Name:  "How to Verify Your Addon",
				Value: "To verify your addon you must create a post in the **verification requests** channel",
			},
			{
				Name:  "What Your Post Must Include",
				Value: "- A GitHub repository link\n- A short description of what your addon does",
			},
			{
				Name:  "What Will Prevent Verification",
				Value: "- Harmful or malicious features (e.g., backdoors, coordinate leaks, remote control)\n- Obfuscated or intentionally confusing code\n- Forks of addons that are still actively maintained (commits within the last 6 months)\n- Copying modules from other addons without proper credit or meaningful changes\n- Addons that appear vibe coded will not be verified for that reason if they have any issues",
			},
			{
				Name:  "",
				Value: "*It may take some time to verify your addon. You will be notified when a decision is made.*\n*Please only submit addons that you have created.*",
			},
			{
				Name:  "Verified Addon Requirements",
				Value: "Once verified, your addon must continue to meet these requirements:\n\n- Code must remain open-source and free of obfuscation\n- No hidden, deceptive, or user-targeting code, including code that:\n  - Disconnects, kicks, or disrupts users\n  - Gives a user remote control over another user's client\n  - Collects or sends private data (coords, IPs, messages) without consent\n  - Cosmetic-only features are allowed\n- Addons should not cause severe performance issues or crashes\n- Addons must stay compatible with current Meteor Client + Minecraft versions",
			},
			{
				Name:  "Policy Violations & Consequences",
				Value: "- **First Offense:** Addon will be unverified until compliant. Developer will be notified if contact info is available.\n- **Second Offense:** Permanently unverified and not eligible for future verification.\n- *Developers may appeal by contacting cqb13 on Discord.*",
			},
			{
				Name:  "Contributing",
				Value: "Help improve Meteor Addon List by contributing to the [website](https://github.com/cqb13/meteor-addons) or [scanner](https://github.com/cqb13/meteor-addon-scanner).",
			},
			{
				Name:  "",
				Value: "*Run **/notified** to stay updated on all new verified addons*",
			},
		},
	}

	if !utils.IsOwner(i) {
		err := utils.InteractionRespondEmbed(embed, s, i.Interaction, true, "")
		if err != nil {
			utils.LogCmdResponseFailiure(i, err.Error())
		}
		return
	}

	err := utils.SendToChannelEmbed(embed, s, i.ChannelID, "")
	if err != nil {
		utils.LogCmdResponseFailiure(i, err.Error())
	}
	err = utils.InteractionRespondText("Sent", s, i.Interaction, true, "")
	if err != nil {
		utils.LogCmdResponseFailiure(i, err.Error())
	}
}

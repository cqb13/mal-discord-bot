package utils

import "github.com/bwmarrin/discordgo"

const OwnerID = "565976415065997312"
const NotifiedRoleId = "1391183408402796687"
const MainServerId = "1377120290898837534"

func IsOwner(i *discordgo.InteractionCreate) bool {
	return i.Member.User.ID == OwnerID
}

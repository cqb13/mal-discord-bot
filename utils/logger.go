package utils

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Log(message string) {
	time := time.Now()
	fmt.Printf("[%02d/%02d/%d | %02d:%02d:%02d] %s\n", time.Month(), time.Day(), time.Year(), time.Hour(), time.Minute(), time.Second(), message)
}

func LogCmdError(i *discordgo.InteractionCreate, error string) {
	Log(fmt.Sprintf("%s command failed internally: %s", i.ApplicationCommandData().Name, error))
}

func LogCmdResponseFailiure(i *discordgo.InteractionCreate, error string) {
	Log(fmt.Sprintf("%s command failed to respond: %s", i.ApplicationCommandData().Name, error))
}

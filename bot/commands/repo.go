package commands

import (
	"dev/cqb13/mal-bot/utils"
	"fmt"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

type repository struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Stars         int      `json:"stargazers_count"`
	Forks         int      `json:"forks_count"`
	DefaultBranch string   `json:"default_branch"`
	PushedAt      string   `json:"pushed_at"`
	CreatedAt     string   `json:"created_at"`
	Fork          bool     `json:"fork"`
	Archived      bool     `json:"archived"`
	Topics        []string `json:"topics"`
	Owner         struct {
		AvatarUrl string `json:"avatar_url"`
	} `json:"owner"`
}

var RepoCommand = &discordgo.ApplicationCommand{
	Name:        "repo",
	Description: "Show information about a github repository",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "owner",
			Description: "Owner of the repository",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "repo",
			Description: "Name of the repository",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func handleRepo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	var owner, repo string

	for _, opt := range options {
		switch opt.Name {
		case "owner":
			owner = opt.StringValue()
		case "repo":
			repo = opt.StringValue()
		}
	}

	repository, err := getRepo(owner, repo)
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	createdAt, err := utils.RFC3339StrToPrettyStr(repository.CreatedAt)
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	pushedAt, err := utils.RFC3339StrToPrettyStr(repository.PushedAt)
	if err != nil {
		utils.InteractionRespondText(fmt.Sprintf("Command Failed: %v", err), s, i.Interaction, true, "")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       repository.Name,
		Description: repository.Description,
		Color:       utils.EmbedColor,
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: repository.Owner.AvatarUrl},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Stars",
				Value:  strconv.Itoa(repository.Stars),
				Inline: true,
			},
			{
				Name:   "Forks",
				Value:  strconv.Itoa(repository.Forks),
				Inline: true,
			},
			{
				Name:   "Default Branch",
				Value:  repository.DefaultBranch,
				Inline: true,
			},
			{
				Name:   "Fork",
				Value:  utils.Ternary(repository.Fork, "True", "False"),
				Inline: true,
			},
			{
				Name:   "Archived",
				Value:  utils.Ternary(repository.Fork, "True", "False"),
				Inline: true,
			},
			{
				Name:   "Created On",
				Value:  createdAt,
				Inline: false,
			},
			{
				Name:   "Latest Push",
				Value:  pushedAt,
				Inline: true,
			},
			{
				Name:  "Topics",
				Value: strings.Join(repository.Topics, ", "),
			},
		},
	}

	utils.InteractionRespondEmbed(embed, s, i.Interaction, false, "")
}

func getRepo(owner string, repo string) (*repository, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	bytes, err := utils.MakeGetRequest(url)
	if err != nil {
		return nil, err
	}

	var repository repository

	err = json.Unmarshal(bytes, &repository)
	if err != nil {
		return nil, err
	}

	return &repository, nil
}

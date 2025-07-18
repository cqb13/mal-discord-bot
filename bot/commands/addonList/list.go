package addonList

import (
	"dev/cqb13/mal-bot/utils"
	"encoding/json"
	"time"
)

type Addon struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	McVersion   string   `json:"mc_version"`
	Authors     []string `json:"authors"`
	Features    Features `json:"features"`
	Verified    bool     `json:"verified"`
	Repo        Repo     `json:"repo"`
	Links       Links    `json:"links"`
}

type Features struct {
	Modules      []string `json:"modules"`
	Commands     []string `json:"commands"`
	HudElements  []string `json:"hud_elements"`
	FeatureCount int      `json:"feature_count"`
}

type Repo struct {
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	Archived     bool   `json:"archived"`
	Fork         bool   `json:"fork"`
	Stars        int    `json:"stars"`
	Downloads    int    `json:"downloads"`
	LastUpdate   string `json:"last_update"`
	CreationDate string `json:"creation_date"`
}

type Links struct {
	Github   string `json:"github"`
	Download string `json:"download"`
	Discord  string `json:"discord"`
	Homepage string `json:"homepage"`
	Icon     string `json:"icon"`
}

var FetchTime time.Time
var currentList []Addon

func UseList() ([]Addon, error) {
	currentTime := time.Now()

	if currentTime.Sub(FetchTime) > time.Hour {
		list, err := getAddonList()
		if err != nil {
			return nil, err
		}
		currentList = list
		FetchTime = currentTime
	}

	return currentList, nil
}

func getAddonList() ([]Addon, error) {
	url := "https://raw.githubusercontent.com/cqb13/meteor-addon-scanner/main/addons.json"
	bytes, err := utils.MakeGetRequest(url)
	if err != nil {
		return nil, err
	}

	var list []Addon
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

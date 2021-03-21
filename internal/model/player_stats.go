package model

type PlayerStats []struct {
	SteamId string `json:"steamID"`

	Stats struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"stats"`

	Achievements struct {
		Name     string `json:"name"`
		Achieved bool   `json:"achieved"`
	} `json:"achievements"`
}

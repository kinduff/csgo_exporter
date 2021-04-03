package model

// PlayerStats stores a struct of JSON response
type PlayerStats []struct {
	SteamID string `json:"steamID"`

	Stats struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"stats"`

	Achievements struct {
		Name     string `json:"name"`
		Achieved bool   `json:"achieved"`
	} `json:"achievements"`
}

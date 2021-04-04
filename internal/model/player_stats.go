package model

// Stats stores a struct of JSON response
type PlayerStats struct {
	PlayerStats struct {
		SteamID string `json:"playerstats.steamID"`

		Stats []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"stats"`

		Achievements []struct {
			Name     string `json:"name"`
			Achieved int    `json:"achieved"`
		} `json:"achievements"`
	} `json:"playerstats"`
}

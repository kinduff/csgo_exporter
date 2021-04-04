// Package which provides structs to be used with the CSGO API
package model

// PlayerStats stores a struct of JSON response for player statistics
// it contains the SteamID which is the ID of the player, Stats which are
// all the data gather from all the game, as well as achievements which
// are the ones achieved by the player.
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

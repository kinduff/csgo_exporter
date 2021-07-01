package model

type PlayerSummaries struct {
	Response struct {
		Players []struct {
			PersonaName string `json:"personaname"`
		} `json:"players"`
	} `json:"response"`
}

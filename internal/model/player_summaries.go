package model

// PlayerSummaries stores information about the player's Steam profile,
// it is used to convert a SteamID to a SteamName. Inverso of ResolveVanityUrl
type PlayerSummaries struct {
	Response struct {
		Players []struct {
			PersonaName string `json:"personaname"`
		} `json:"players"`
	} `json:"response"`
}

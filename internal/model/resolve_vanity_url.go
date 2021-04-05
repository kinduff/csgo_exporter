package model

// ResolveVanityUrl stores information about the player's Steam ID,
// it is used to convert a username to a SteamID that can be used
// for future requests.
type ResolveVanityUrl struct {
	Response struct {
		Steamid string `json:"steamid"`
	} `json:"response"`
}

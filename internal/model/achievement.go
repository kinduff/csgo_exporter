package model

// Achievement stores the complete information about achievements of a player.
// this includes metadata to identify what and how, as well if it was achieved
// by the player. This struct is filled in combination of a Stats request, and
// the parsing of an XML request from the user's public profile.
type Achievement struct {
	APIName     string
	Title       string
	Description string
	Achieved    int
}

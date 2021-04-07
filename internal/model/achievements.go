package model

// Achievements stores a struct of JSON response for global achievements percentage,
// this one is used to see all the achievements from a game, since the calls from
// PlayerStats doesn't return the ones you don't have.
type Achievements struct {
	AchievementPercentages struct {
		Achievements []struct {
			Name    string  `json:"name"`
			Percent float64 `json:"percent"`
		} `json:"achievements"`
	} `json:"achievementpercentages"`
}

package model

// AchievementsDetails stores a struct of a XML response to complement the response
// from the API, since the other response only contains raw data, this adds more detail.
type AchievementsDetails struct {
	Achievements struct {
		Achievement []struct {
			Closed      string `xml:"closed,attr"`
			Name        string `xml:"name"`
			Apiname     string `xml:"apiname"`
			Description string `xml:"description"`
		} `xml:"achievement"`
	} `xml:"achievements"`
}

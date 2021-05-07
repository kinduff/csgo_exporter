package model

import (
	"encoding/xml"
)

// AchievementsDetails stores a struct of a XML response to complement the response
// from the API, since the other response only contains raw data, this adds more detail.
type AchievementsDetails struct {
	XMLName         xml.Name `xml:"playerstats"`
	Text            string   `xml:",chardata"`
	PrivacyState    string   `xml:"privacyState"`
	VisibilityState string   `xml:"visibilityState"`
	Game            struct {
		Text             string `xml:",chardata"`
		GameFriendlyName string `xml:"gameFriendlyName"`
		GameName         string `xml:"gameName"`
		GameLink         string `xml:"gameLink"`
		GameIcon         string `xml:"gameIcon"`
		GameLogo         string `xml:"gameLogo"`
		GameLogoSmall    string `xml:"gameLogoSmall"`
	} `xml:"game"`
	Player struct {
		Text      string `xml:",chardata"`
		SteamID64 string `xml:"steamID64"`
		CustomURL string `xml:"customURL"`
	} `xml:"player"`
	Stats struct {
		Text        string `xml:",chardata"`
		HoursPlayed string `xml:"hoursPlayed"`
	} `xml:"stats"`
	Achievements struct {
		Text        string `xml:",chardata"`
		Achievement []struct {
			Text            string `xml:",chardata"`
			Closed          string `xml:"closed,attr"`
			IconClosed      string `xml:"iconClosed"`
			IconOpen        string `xml:"iconOpen"`
			Name            string `xml:"name"`
			Apiname         string `xml:"apiname"`
			Description     string `xml:"description"`
			UnlockTimestamp string `xml:"unlockTimestamp"`
		} `xml:"achievement"`
	} `xml:"achievements"`
}

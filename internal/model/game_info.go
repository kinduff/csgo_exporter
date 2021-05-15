package model

// GameInfo stores information about the game itself from a player perspective,
// this includes the playtime in different categories including per OS.
type GameInfo struct {
	Response struct {
		Games []struct {
			Playtime2Weeks         int `json:"playtime_2weeks"`
			PlaytimeForever        int `json:"playtime_forever"`
			PlaytimeWindowsForever int `json:"playtime_windows_forever"`
			PlaytimeMacForever     int `json:"playtime_mac_forever"`
			PlaytimeLinuxForever   int `json:"playtime_linux_forever"`
		} `json:"games"`
	} `json:"response"`
}

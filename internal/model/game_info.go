package model

type GameInfo struct {
	Response struct {
		Games []struct {
			Appid                  int `json:"appid"`
			Playtime2Weeks         int `json:"playtime_2weeks"`
			PlaytimeForever        int `json:"playtime_forever"`
			PlaytimeWindowsForever int `json:"playtime_windows_forever"`
			PlaytimeMacForever     int `json:"playtime_mac_forever"`
			PlaytimeLinuxForever   int `json:"playtime_linux_forever"`
		} `json:"games"`
	} `json:"response"`
}

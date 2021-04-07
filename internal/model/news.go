package model

// News stores a struct of JSON response for news from CS:GO.
type News struct {
	Appnews struct {
		Appid     int `json:"appid"`
		Newsitems []struct {
			Gid           string `json:"gid"`
			Title         string `json:"title"`
			URL           string `json:"url"`
			IsExternalURL bool   `json:"is_external_url"`
			Author        string `json:"author"`
			Contents      string `json:"contents"`
			Feedlabel     string `json:"feedlabel"`
			Date          int    `json:"date"`
			Feedname      string `json:"feedname"`
			FeedType      int    `json:"feed_type"`
			Appid         int    `json:"appid"`
		} `json:"newsitems"`
		Count int `json:"count"`
	} `json:"appnews"`
}

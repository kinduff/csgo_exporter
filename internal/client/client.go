package client

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	httpClient http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (client *Client) DoRequest(steamID string, apiKey string, target interface{}) error {
	req, err := http.NewRequest("GET", "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002", nil)
	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)
		return err
	}

	q := req.URL.Query()
	q.Add("appid", "730")
	q.Add("steamid", steamID)
	q.Add("key", apiKey)

	req.URL.RawQuery = q.Encode()

	log.Infof("Sending HTTP request to %s", req.URL.String())

	resp, err := client.httpClient.Do(req)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Fatalf("An error has occurred during retrieving statistics %v", err)
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

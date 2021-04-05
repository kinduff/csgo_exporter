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

func getEndpoint(endpoint string) string {
	var path string
	baseUrl := "https://api.steampowered.com/"

	switch endpoint {
	case "stats":
		path = "ISteamUserStats/GetUserStatsForGame/v0002"
	case "id":
		path = "ISteamUser/ResolveVanityURL/v0001"
	}

	return baseUrl + path
}

func getQueryParams(req *http.Request, apiKey string, endpoint string, steamID string) string {
	q := req.URL.Query()
	q.Add("appid", "730")
	q.Add("key", apiKey)

	switch endpoint {
	case "stats":
		q.Add("steamid", steamID)
	case "id":
		q.Add("vanityurl", steamID)
	}

	return q.Encode()
}

func (client *Client) DoRequest(endpoint string, steamID string, apiKey string, target interface{}) error {
	req, err := http.NewRequest("GET", getEndpoint(endpoint), nil)
	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)

		return err
	}

	req.URL.RawQuery = getQueryParams(req, apiKey, endpoint, steamID)

	log.Infof("Sending HTTP request to %s", req.URL.String())

	resp, err := client.httpClient.Do(req)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Fatalf("An error has occurred during retrieving statistics %v", err)

		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

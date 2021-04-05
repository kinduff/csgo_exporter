package client

import (
	"encoding/json"
	"net/http"

	"github.com/kinduff/csgo_exporter/internal/model"

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
	baseUrl := "https://api.steampowered.com"

	switch endpoint {
	case "stats":
		path = "/ISteamUserStats/GetUserStatsForGame/v0002"
	case "id":
		path = "/ISteamUser/ResolveVanityURL/v0001"
	}

	return baseUrl + path
}

func getQueryParams(endpoint string, config *model.Config, req *http.Request) string {
	q := req.URL.Query()
	q.Add("appid", "730")
	q.Add("key", config.ApiKey)

	switch endpoint {
	case "stats":
		q.Add("steamid", config.SteamID)
	case "id":
		q.Add("vanityurl", config.SteamName)
	}

	return q.Encode()
}

func (client *Client) DoRequest(endpoint string, config *model.Config, target interface{}) error {
	req, err := http.NewRequest("GET", getEndpoint(endpoint), nil)
	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)

		return err
	}

	req.URL.RawQuery = getQueryParams(endpoint, config, req)

	log.Infof("Sending HTTP request to %s", req.URL.String())

	resp, err := client.httpClient.Do(req)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Fatalf("An error has occurred during retrieving statistics %v", err)

		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

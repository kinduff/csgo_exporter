package client

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kinduff/csgo_exporter/internal/model"

	log "github.com/sirupsen/logrus"
)

// Client is a struct that contains an HTTP Client
type Client struct {
	httpClient http.Client
}

// NewClient provides an interface to make HTTP requests to the Steam API.
func NewClient() *Client {
	return &Client{
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// DoAPIRequest allows to make requests to the Steam API by standarizing how it receives
// parameters, and to which endpoint it should do the call.
func (client *Client) DoAPIRequest(endpoint string, config *model.Config, target interface{}) error {
	req, err := http.NewRequest("GET", getAPIEndpoint(endpoint), nil)
	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)

		return err
	}

	req.URL.RawQuery = getAPIQueryParams(endpoint, config, req)

	log.Infof("Sending HTTP request to %s", req.URL.String())

	resp, err := client.httpClient.Do(req)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Fatalf("An error has occurred during retrieving statistics %v", err)

		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func getAPIEndpoint(endpoint string) string {
	var path string
	baseUrl := "https://api.steampowered.com"

	switch endpoint {
	case "achievements":
		path = "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002"
	case "stats":
		path = "/ISteamUserStats/GetUserStatsForGame/v0002"
	case "id":
		path = "/ISteamUser/ResolveVanityURL/v0001"
	case "news":
		path = "/ISteamNews/GetNewsForApp/v0002"
	case "gameInfo":
		path = "/IPlayerService/GetOwnedGames/v1"
	}

	return baseUrl + path
}

func getAPIQueryParams(endpoint string, config *model.Config, req *http.Request) string {
	q := req.URL.Query()
	q.Add("key", config.ApiKey)

	gameIdKey := "appid"

	switch endpoint {
	case "achievements":
		gameIdKey = "gameid"
	case "stats":
		q.Add("steamid", config.SteamID)
	case "id":
		q.Add("vanityurl", config.SteamName)
	case "news":
		q.Add("maxlength", "240")
	case "gameInfo":
		q.Add("steamid", config.SteamID)
		q.Add("appids_filter[0]", "730")
	}

	q.Add(gameIdKey, "730")

	return q.Encode()
}

// DoXMLRequest allows to make requests to the Steam web API, this API is not documented,
// and it's a hacky way to access certain data in XML. It relies on the user having its profile public
func (client *Client) DoXMLRequest(endpoint string, config *model.Config, target interface{}) error {
	req, err := http.NewRequest("GET", getXMLEndpoint(endpoint, config), nil)
	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)

		return err
	}

	log.Infof("Sending HTTP request to %s", req.URL.String())

	resp, err := client.httpClient.Do(req)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Fatalf("An error has occurred during retrieving statistics %v", err)

		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return xml.Unmarshal(data, &target)
}

func getXMLEndpoint(endpoint string, config *model.Config) string {
	var path string
	baseUrl := fmt.Sprintf("https://steamcommunity.com/profiles/%s", config.SteamID)

	switch endpoint {
	case "achievementsDetails":
		path = "/stats/CSGO?xml=1"
	}

	return baseUrl + path
}

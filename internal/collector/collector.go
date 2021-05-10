package collector

import (
	"strings"
	"time"

	"github.com/kinduff/csgo_exporter/internal/client"
	"github.com/kinduff/csgo_exporter/internal/metrics"
	"github.com/kinduff/csgo_exporter/internal/model"

	log "github.com/sirupsen/logrus"
)

type collector struct {
	config *model.Config
}

// NewCollector provides an interface to collector player statistics.
func NewCollector(config *model.Config) *collector {
	return &collector{
		config: config,
	}
}

func (collector *collector) Scrape() {
	for range time.Tick(collector.config.ScrapeInterval) {
		collector.setMetrics()
		log.Printf("New tick of statistics")
	}
}

func (collector *collector) setMetrics() {
	client := client.NewClient()

	if collector.config.SteamID == "" {
		ResolveVanityUrl := model.ResolveVanityUrl{}
		if err := client.DoAPIRequest("id", collector.config, &ResolveVanityUrl); err != nil {
			log.Fatal(err)
		}
		collector.config.SteamID = ResolveVanityUrl.Response.Steamid
	}

	player := collector.config.SteamName
	if player == "" {
		player = collector.config.SteamID
	}

	playerStats := model.PlayerStats{}
	if err := client.DoAPIRequest("stats", collector.config, &playerStats); err != nil {
		log.Fatal(err)
	}

	news := model.News{}
	if err := client.DoAPIRequest("news", collector.config, &news); err != nil {
		log.Fatal(err)
	}

	gameInfo := model.GameInfo{}
	if err := client.DoAPIRequest("gameInfo", collector.config, &gameInfo); err != nil {
		log.Fatal(err)
	}

	achievementsDetails := model.AchievementsDetails{}
	if err := client.DoXMLRequest("achievementsDetails", collector.config, &achievementsDetails); err != nil {
		log.Fatal(err)
	}

	allPlayerAchievements := map[string]model.Achievement{}

	for _, s := range achievementsDetails.Achievements.Achievement {
		allPlayerAchievements[s.APIName] = model.Achievement{
			APIName:     s.APIName,
			Achieved:    0,
			Title:       s.Name,
			Description: s.Description,
		}
	}

	playerAchievements := playerStats.PlayerStats.Achievements
	for _, s := range playerAchievements {
		t := allPlayerAchievements[strings.ToLower(s.Name)]

		allPlayerAchievements[t.APIName] = model.Achievement{
			APIName:     t.APIName,
			Achieved:    1,
			Title:       t.Title,
			Description: t.Description,
		}
	}

	for _, s := range playerStats.PlayerStats.Stats {
		if strings.Contains(s.Name, "GI") {
			continue
		}

		metrics.Stats.WithLabelValues(player, s.Name).Set(float64(s.Value))
	}

	for _, s := range allPlayerAchievements {
		metrics.Achievements.WithLabelValues(player, s.APIName, s.Title, s.Description).Set(float64(s.Achieved))
	}

	playData := gameInfo.Response.Games[0]
	metrics.Playtime.WithLabelValues(player, "last_2_weeks").Set(float64(playData.Playtime2Weeks))
	metrics.Playtime.WithLabelValues(player, "forever").Set(float64(playData.PlaytimeForever))
	metrics.Playtime.WithLabelValues(player, "windows_forever").Set(float64(playData.PlaytimeWindowsForever))
	metrics.Playtime.WithLabelValues(player, "mac_forever").Set(float64(playData.PlaytimeMacForever))
	metrics.Playtime.WithLabelValues(player, "linux_forever").Set(float64(playData.PlaytimeLinuxForever))

	for _, s := range news.Appnews.Newsitems {
		metrics.News.WithLabelValues(player, s.Title, s.URL, s.Feedlabel).Set(float64(s.Date) * 1000)
	}
}

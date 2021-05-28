// Package collector handles the orchestration between the API and Prometheus.
package collector

import (
	"time"

	"github.com/kinduff/csgo_exporter/config"
	"github.com/kinduff/csgo_exporter/internal/client"
	"github.com/kinduff/csgo_exporter/internal/model"
)

type collector struct {
	config                *config.Config
	playerStats           model.PlayerStats
	news                  model.News
	gameInfo              model.GameInfo
	playerInventory       map[string]model.PlayerInventory
	allPlayerAchievements map[string]model.Achievement
	client                *client.Client
}

// NewCollector provides an interface to collector player statistics.
func NewCollector(config *config.Config) *collector {
	return &collector{
		config:                config,
		playerStats:           model.PlayerStats{},
		news:                  model.News{},
		gameInfo:              model.GameInfo{},
		playerInventory:       map[string]model.PlayerInventory{},
		allPlayerAchievements: map[string]model.Achievement{},
		client:                client.NewClient(),
	}
}

func (collector *collector) Scrape() {
	for range time.Tick(collector.config.ScrapeInterval) {
		if collector.config.SteamID == "" {
			collector.collectSteamID()
		}
		go collector.collectStats()
		go collector.collectAchievements()
		go collector.collectGameInfo()
		go collector.collectNews()
		go collector.collectPlayerInventory()
	}
}

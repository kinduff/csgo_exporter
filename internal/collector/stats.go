package collector

import (
	"strings"

	"github.com/kinduff/csgo_exporter/internal/metrics"
	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectStats() {
	if err := collector.client.DoAPIRequest("stats", collector.config, &collector.playerStats); err != nil {
		log.Fatal(err)
	}

	for _, s := range collector.playerStats.PlayerStats.Stats {
		if strings.Contains(s.Name, "GI") {
			continue
		}

		metrics.Stats.WithLabelValues(collector.config.SteamID, s.Name).Set(float64(s.Value))

		if strings.Contains(s.Name, "last_match") {
			metrics.LastMatch.WithLabelValues(collector.config.SteamID, strings.Split(s.Name, "last_match_")[1]).Set(float64(s.Value))
		}

		if strings.Contains(s.Name, "total_shots") {
			metrics.TotalShots.WithLabelValues(collector.config.SteamID, strings.Split(s.Name, "total_shots_")[1]).Set(float64(s.Value))
		}

		if strings.Contains(s.Name, "total_kills_") {
			weapon_name := strings.Split(s.Name, "total_kills_")[1]
			excluded_names := []string{"against_zoomed_sniper", "enemy_blinded", "enemy_weapon", "knife_fight"}

			found := find(excluded_names, weapon_name)

			if !found {
				metrics.TotalKills.WithLabelValues(collector.config.SteamID, weapon_name).Set(float64(s.Value))
			}
		}
	}
}

// find takes a slice and looks for an element in it. If found it will
// return true, otherwise it will return false.
// source: https://golangcode.com/check-if-element-exists-in-slice
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

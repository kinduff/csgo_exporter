package collector

import (
	"strings"

	"github.com/kinduff/csgo_exporter/internal/data"
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

		metrics.Stats.WithLabelValues(collector.config.SteamID, collector.config.SteamName, s.Name).Set(float64(s.Value))

		if strings.Contains(s.Name, "last_match") {
			title := ""
			statName := strings.Split(s.Name, "last_match_")[1]
			if statName == "favweapon_id" {
				title = data.WeaponByID(s.Value)
			}
			metrics.LastMatch.WithLabelValues(collector.config.SteamID, collector.config.SteamName, statName, title).Set(float64(s.Value))
		}

		if strings.Contains(s.Name, "total_shots") {
			title := strings.Split(s.Name, "total_shots_")[1]
			if title != "fired" && title != "hit" {
				title = data.WeaponByAPIName(title)
				metrics.TotalShots.WithLabelValues(collector.config.SteamID, collector.config.SteamName, title).Set(float64(s.Value))
			}
		}

		if strings.Contains(s.Name, "total_hits") {
			title := strings.Split(s.Name, "total_hits_")[1]
			title = data.WeaponByAPIName(title)
			metrics.TotalHits.WithLabelValues(collector.config.SteamID, collector.config.SteamName, title).Set(float64(s.Value))
		}

		if strings.Contains(s.Name, "total_kills_") {
			weaponName := strings.Split(s.Name, "total_kills_")[1]
			excludedNames := []string{"against_zoomed_sniper", "enemy_blinded", "enemy_weapon", "knife_fight", "headshot"}

			found := find(excludedNames, weaponName)

			if !found {
				title := data.WeaponByAPIName(weaponName)
				metrics.TotalKills.WithLabelValues(collector.config.SteamID, collector.config.SteamName, title).Set(float64(s.Value))
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

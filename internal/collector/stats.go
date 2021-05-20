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
	}

}

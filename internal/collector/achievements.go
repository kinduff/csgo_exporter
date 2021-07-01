package collector

import (
	"strings"

	"github.com/kinduff/csgo_exporter/internal/metrics"
	"github.com/kinduff/csgo_exporter/internal/model"
	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectAchievements() {
	achievementsDetails := model.AchievementsDetails{}
	if err := collector.client.DoXMLRequest("achievementsDetails", collector.config, &achievementsDetails); err != nil {
		log.Fatal(err)
	}

	for _, s := range achievementsDetails.Achievements.Achievement {
		collector.allPlayerAchievements[s.APIName] = model.Achievement{
			APIName:     s.APIName,
			Achieved:    0,
			Title:       s.Name,
			Description: s.Description,
		}
	}

	for _, s := range collector.playerStats.PlayerStats.Achievements {
		t := collector.allPlayerAchievements[strings.ToLower(s.Name)]

		collector.allPlayerAchievements[t.APIName] = model.Achievement{
			APIName:     t.APIName,
			Achieved:    1,
			Title:       t.Title,
			Description: t.Description,
		}
	}

	for _, s := range collector.allPlayerAchievements {
		metrics.Achievements.WithLabelValues(collector.config.SteamID, collector.config.SteamName, s.APIName, s.Title, s.Description).Set(float64(s.Achieved))
	}
}

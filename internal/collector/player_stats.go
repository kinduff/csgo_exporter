package collector

import (
	"github.com/kinduff/csgo_exporter/internal/client"
	"github.com/kinduff/csgo_exporter/internal/model"

	"github.com/prometheus/client_golang/prometheus"

	log "github.com/sirupsen/logrus"
)

type playerCollector struct {
	steamID            string
	apiKey             string
	statsMetric        *prometheus.Desc
	achievementsMetric *prometheus.Desc
}

func NewPlayerCollector(steamID string, apiKey string) *playerCollector {
	return &playerCollector{
		steamID: steamID,
		apiKey:  apiKey,
		statsMetric: prometheus.NewDesc("stats_metric",
			"Shows metrics a player has from all its matches",
			[]string{"name"},
			prometheus.Labels{"steamID": steamID},
		),
		achievementsMetric: prometheus.NewDesc("achievements_metric",
			"Shows metrics a player has for its achievements",
			[]string{"name"},
			prometheus.Labels{"steamID": steamID},
		),
	}
}

func (collector *playerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.statsMetric
	ch <- collector.achievementsMetric
}

func (collector *playerCollector) Collect(ch chan<- prometheus.Metric) {
	c := client.NewClient()

	ResolveVanityUrl := model.ResolveVanityUrl{}
	if err := c.DoRequest("id", collector.steamID, collector.apiKey, &ResolveVanityUrl); err != nil {
		log.Fatal(err)
	}

	playerStats := model.PlayerStats{}
	if err := c.DoRequest("stats", ResolveVanityUrl.Response.Steamid, collector.apiKey, &playerStats); err != nil {
		log.Fatal(err)
	}

	stats := playerStats.PlayerStats.Stats
	for _, s := range stats {
		ch <- prometheus.MustNewConstMetric(collector.statsMetric, prometheus.GaugeValue, float64(s.Value), s.Name)
	}

	achievements := playerStats.PlayerStats.Achievements
	for _, s := range achievements {
		ch <- prometheus.MustNewConstMetric(collector.achievementsMetric, prometheus.GaugeValue, float64(s.Achieved), s.Name)
	}
}

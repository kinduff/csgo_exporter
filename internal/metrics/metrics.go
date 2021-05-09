package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	log "github.com/sirupsen/logrus"
)

var (
	namespace = "csgo"

	Stats = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "stats_metric",
			Help:      "Shows metrics a player has from all its matches",
			Namespace: namespace,
		},
		[]string{"player", "name"},
	)
	Achievements = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "achievements_metric",
			Help:      "Shows all the achievements from a player",
			Namespace: namespace,
		},
		[]string{"player", "name", "title", "description"},
	)
	Playtime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "playtime_metric",
			Help:      "Shows the playtime the user has in the game in minutes",
			Namespace: namespace,
		},
		[]string{"player", "type"},
	)
	News = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "news_metric",
			Help:      "Shows the latest news from CSGO",
			Namespace: namespace,
		},
		[]string{"player", "title", "url", "feedlabel"},
	)
)

// Init initializes all Prometheus metrics
func Init() {
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	initMetric("stats", Stats)
	initMetric("achievements", Achievements)
	initMetric("playtime", Playtime)
	initMetric("news", News)
}

func initMetric(name string, metric *prometheus.GaugeVec) {
	prometheus.MustRegister(metric)
	log.Printf("New Prometheus metric registered: %s", name)
}

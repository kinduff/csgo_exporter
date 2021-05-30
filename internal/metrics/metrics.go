// Package metrics sets and initializes Prometheus metrics.
package metrics

import (
	"github.com/kinduff/csgo_exporter/config"
	"github.com/prometheus/client_golang/prometheus"

	log "github.com/sirupsen/logrus"
)

var (
	namespace = "csgo"

	// Stats - Miscellaneous stats of a player from all its matches.
	Stats = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "stats_metric",
			Help:      "Shows metrics a player has from all its matches",
			Namespace: namespace,
		},
		[]string{"player", "name"},
	)

	// LastMatch - Stats of a player from its last match.
	LastMatch = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "last_match_metric",
			Help:      "Shows metrics from a player last match",
			Namespace: namespace,
		},
		[]string{"player", "type"},
	)

	// TotalShots - Total shots per weapon.
	TotalShots = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "total_shots_metric",
			Help:      "Shows total shots from a player per weapon ",
			Namespace: namespace,
		},
		[]string{"player", "name"},
	)

	// TotalKills - Total kills per weapon.
	TotalKills = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "total_kills_metric",
			Help:      "Shows total kills from a player per weapon ",
			Namespace: namespace,
		},
		[]string{"player", "name"},
	)

	// Achievements - All achievements a player can have, including the ones it has.
	Achievements = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "achievements_metric",
			Help:      "Shows all the achievements from a player",
			Namespace: namespace,
		},
		[]string{"player", "name", "title", "description"},
	)

	// Playtime - Hours spent playing the game in minutes from different OS.
	Playtime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "playtime_metric",
			Help:      "Shows the playtime the user has in the game in minutes",
			Namespace: namespace,
		},
		[]string{"player", "type"},
	)

	// News - The latest news from the CSGO community and Valve.
	News = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "news_metric",
			Help:      "Shows the latest news from CSGO",
			Namespace: namespace,
		},
		[]string{"player", "title", "url", "feedlabel"},
	)

	// UserInventory - User inventory with the value as the average price.
	UserInventory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "user_inventory_metric",
			Help:      "Shows the content of the users inventory",
			Namespace: namespace,
		},
		[]string{"player", "class_id", "market_name", "currency", "amount", "tradable", "marketable"},
	)
)

// Init initializes all Prometheus metrics
func Init(config *config.Config) {
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	initMetric("stats", Stats)
	initMetric("last_match", LastMatch)
	initMetric("total_shots", TotalShots)
	initMetric("total_kills", TotalKills)
	initMetric("achievements", Achievements)
	initMetric("playtime", Playtime)
	initMetric("news", News)

	if config.FetchInventory {
		initMetric("user_inventory", UserInventory)
	}
}

func initMetric(name string, metric *prometheus.GaugeVec) {
	prometheus.MustRegister(metric)
	log.Printf("New Prometheus metric registered: %s", name)
}

package collector

import (
	"github.com/kinduff/csgo_exporter/internal/metrics"
	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectNews() {
	if err := collector.client.DoAPIRequest("news", collector.config, &collector.news); err != nil {
		log.Fatal(err)
	}

	for _, s := range collector.news.Appnews.Newsitems {
		metrics.News.WithLabelValues(collector.config.SteamID, s.Title, s.URL, s.Feedlabel).Set(float64(s.Date) * 1000)
	}
}

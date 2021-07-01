package collector

import (
	"github.com/kinduff/csgo_exporter/internal/metrics"
	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectGameInfo() {
	if err := collector.client.DoAPIRequest("gameInfo", collector.config, &collector.gameInfo); err != nil {
		log.Fatal(err)
	}

	playData := collector.gameInfo.Response.Games[0]
	metrics.Playtime.WithLabelValues(collector.config.SteamID, collector.config.SteamName, "last_2_weeks").Set(float64(playData.Playtime2Weeks))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, collector.config.SteamName, "forever").Set(float64(playData.PlaytimeForever))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, collector.config.SteamName, "windows_forever").Set(float64(playData.PlaytimeWindowsForever))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, collector.config.SteamName, "mac_forever").Set(float64(playData.PlaytimeMacForever))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, collector.config.SteamName, "linux_forever").Set(float64(playData.PlaytimeLinuxForever))
}

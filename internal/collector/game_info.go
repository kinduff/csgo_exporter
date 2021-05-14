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
	metrics.Playtime.WithLabelValues(collector.config.SteamID, "last_2_weeks").Set(float64(playData.Playtime2Weeks))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, "forever").Set(float64(playData.PlaytimeForever))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, "windows_forever").Set(float64(playData.PlaytimeWindowsForever))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, "mac_forever").Set(float64(playData.PlaytimeMacForever))
	metrics.Playtime.WithLabelValues(collector.config.SteamID, "linux_forever").Set(float64(playData.PlaytimeLinuxForever))
}

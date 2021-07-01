package collector

import (
	"github.com/kinduff/csgo_exporter/internal/model"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectSteamName() {
	PlayerSummaries := model.PlayerSummaries{}
	if err := collector.client.DoAPIRequest("name", collector.config, &PlayerSummaries); err != nil {
		log.Fatal(err)
	}
	collector.config.SteamName = PlayerSummaries.Response.Players[0].PersonaName
}

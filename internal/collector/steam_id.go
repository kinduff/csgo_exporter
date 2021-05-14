package collector

import (
	"github.com/kinduff/csgo_exporter/internal/model"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectSteamID() {
	ResolveVanityUrl := model.ResolveVanityUrl{}
	if err := collector.client.DoAPIRequest("id", collector.config, &ResolveVanityUrl); err != nil {
		log.Fatal(err)
	}
	collector.config.SteamID = ResolveVanityUrl.Response.Steamid
}

package collector

import (
	"fmt"
	"strconv"

	"github.com/kinduff/csgo_exporter/internal/metrics"
	"github.com/kinduff/csgo_exporter/internal/model"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectPlayerInventory() {
	inventory := model.Inventory{}
	inventoryEndpoint := fmt.Sprintf("https://steamcommunity.com/inventory/%s/730/2", collector.config.SteamID)
	if err := collector.client.DoCustomAPIRequest(inventoryEndpoint, collector.config, &inventory); err != nil {
		log.Fatal(err)
	}

	itemsList := model.ItemsList{}
	pricesEndpoint := "https://csgobackpack.net/api/GetItemsList/v2/?no_details=true"
	if err := collector.client.DoCustomAPIRequest(pricesEndpoint, collector.config, &itemsList); err != nil {
		log.Fatal(err)
	}

	for _, s := range inventory.Assets {
		amount, _ := strconv.ParseInt(s.Amount, 10, 64)

		collector.playerInventory[s.ClassID] = model.PlayerInventory{
			ClassID: s.ClassID,
			Amount:  amount,
		}
	}

	for _, s := range inventory.Descriptions {
		asset := collector.playerInventory[s.ClassID]
		market := itemsList.ItemsList[s.MarketName]

		collector.playerInventory[asset.ClassID] = model.PlayerInventory{
			ClassID:      asset.ClassID,
			Currency:     itemsList.Currency,
			Amount:       asset.Amount,
			Tradable:     s.Tradable == 1,
			Marketable:   s.Marketable == 1,
			MarketName:   s.MarketName,
			AveragePrice: market.Price.SevenDays.Average,
		}
	}

	for _, s := range collector.playerInventory {
		metrics.UserInventory.WithLabelValues(
			collector.config.SteamID,
			s.ClassID,
			s.MarketName,
			s.Currency,
			strconv.FormatInt(s.Amount, 16),
			strconv.FormatBool(s.Tradable),
			strconv.FormatBool(s.Marketable),
		).Set(s.AveragePrice)
	}
}

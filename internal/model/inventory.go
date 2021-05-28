package model

// Inventory stores a struct of JSON response for an unofficial API of CS:GO.
type Inventory struct {
	Assets []struct {
		ClassID    string `json:"classid"`
		InstanceID string `json:"instanceid"`
		Amount     string `json:"amount"`
	} `json:"assets"`
	Descriptions []struct {
		ClassID                   string `json:"classid"`
		InstanceID                string `json:"instanceid"`
		Tradable                  int    `json:"tradable"`
		Name                      string `json:"name"`
		Type                      string `json:"type"`
		MarketName                string `json:"market_name"`
		Commodity                 int    `json:"commodity"`
		MarketTradableRestriction int    `json:"market_tradable_restriction"`
		Marketable                int    `json:"marketable"`
	} `json:"descriptions"`
	TotalInventoryCount int `json:"total_inventory_count"`
}

package model

// PlayerInventory stores a struct of all the player's inventory from CS:GO.
type PlayerInventory struct {
	ClassID      string
	MarketName   string
	Amount       int64
	Currency     string
	AveragePrice float64 // decimal
	Tradable     bool
	Marketable   bool
}

package model

// ItemsList returns a struct from CSGOBackpack for all the prices available.
type ItemsList struct {
	Success   bool   `json:"success"`
	Currency  string `json:"currency"`
	Timestamp int    `json:"timestamp"`
	ItemsList map[string]struct {
		Name  string `json:"name"`
		Price struct {
			SevenDays struct {
				Average float64 `json:"average"`
			} `json:"7_days"`
		} `json:"price"`
		FirstSaleDate string `json:"first_sale_date"`
	} `json:"items_list"`
}

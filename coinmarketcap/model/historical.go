package model

type CryptoHistoricalData struct {
	Data map[string][]*CryptoHistorical `json:"data"`
}

type CryptoHistorical struct {
	ID       int                 `json:"id,omitempty"`
	Name     string              `json:"name"`
	Symbol   string              `json:"symbol"`
	IsActive int                 `json:"is_active"`
	IsFiat   int                 `json:"is_fiat"`
	Quotes   []*HistoricalPrices `json:"quotes"`
}

type HistoricalPrices struct {
	Timestamp string `json:"timestamp"`
	Quote     *USD   `json:"quote"`
}

type USD struct {
	USD *HistoricalPriceSnapshot `json:"USD"`
}

type HistoricalPriceSnapshot struct {
	Price             float64 `json:"price"`
	Volume24h         float64 `json:"volume_24h"`
	MarketCap         float64 `json:"market_cap"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	Timestamp         string  `json:"timestamp"`
}

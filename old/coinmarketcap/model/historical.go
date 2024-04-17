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
	Timestamp string         `json:"timestamp"`
	Quote     *USDHistorical `json:"quote"`
}

type USDHistorical struct {
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

type OHLCVHistoricalData struct {
	Data *OHLCVHistorical `json:"data"`
}

type OHLCVHistorical struct {
	ID     int                        `json:"id"`
	Name   string                     `json:"name"`
	Symbol string                     `json:"symbol"`
	Quotes []*OHLCVHistoricalSnapshot `json:"quotes"`
}

type OHLCVHistoricalSnapshot struct {
	TimeOpen  string              `json:"time_open"`
	TimeClose string              `json:"time_close"`
	TimeHigh  string              `json:"time_high"`
	TimeLow   string              `json:"time_low"`
	Quote     *OHLCVHistoricalUSD `json:"quote"`
}

type OHLCVHistoricalUSD struct {
	USD struct {
		Open      float64 `json:"open"`
		High      float64 `json:"high"`
		Low       float64 `json:"low"`
		Close     float64 `json:"close"`
		Timestamp string  `json:"timestamp"`
	} `json:"USD"`
}

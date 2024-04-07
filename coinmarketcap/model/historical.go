package model

import (
	"context"
	"net/url"
)

type QuoteHistoryer interface {
	Get(ctx context.Context, query *url.Values) (*Crypto, error)
}

type HistoricalData struct {
	Data *Crypto `json:"data"`
}

type CoinHistorical struct {
	ID       int                 `json:"id,omitempty"`
	Name     string              `json:"name"`
	Symbol   string              `json:"symbol"`
	IsActive int                 `json:"is_active"`
	IsFiat   int                 `json:"is_fiat"`
	Quotes   []*HistoricalQuotes `json:"quotes"`
}

type Crypto struct {
	CORE []*CoinHistorical `json:"core"`
}

type HistoricalQuotes struct {
	Timestamp string `json:"timestamp"`
	Quote     *USD   `json:"quote"`
}

type USD struct {
	USD *USDQuoteHistorical `json:"USD"`
}

type USDQuoteHistorical struct {
	Price             float64 `json:"price"`
	Volume24h         float64 `json:"volume_24h"`
	MarketCap         float64 `json:"market_cap"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	Timestamp         string  `json:"timestamp"`
}

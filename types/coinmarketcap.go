package types

type Coin struct {
	ID                int       `json:"id,omitempty"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	CMCRank           int       `json:"cmc_rank"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	MaxSupply         float64   `json:"max_supply,omitempty"`
	Tags              []string  `json:"tags"`
	Platform          *Platform `json:"platform"`
	Quote             *Quote    `json:"quote"`
	LastUpdated       string    `json:"last_updated"`
}

type Platform struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	TokenAddress string `json:"token_address"`
}

type Quote struct {
	USD *USD `json:"USD"`
}

type USD struct {
	MarketCap          string  `json:"market_cap"`
	Price              float64 `json:"price"`
	MarketCapDominance float64 `json:"market_cap_dominance"`
	PercentChange1H    float64 `json:"percent_change_1h"`
	PercentChange24H   float64 `json:"percent_change_24h"`
	PercentChange7D    float64 `json:"percent_change_7d"`
	PercentChange30D   float64 `json:"percent_change_30d"`
	Volume24H          float64 `json:"volume_24h"`
	Volume7D           float64 `json:"volume_7d,omitempty"`
	Volume30D          float64 `json:"volume_30d,omitempty"`
	LastUpdated        string  `json:"last_updated"`
}

package model

type Platform struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	TokenAddress string `json:"token_address"`
}

// Listing Latest

type Listing struct {
	ID                int          `json:"id,omitempty"`
	Name              string       `json:"name"`
	Symbol            string       `json:"symbol"`
	Slug              string       `json:"slug"`
	CMCRank           int          `json:"cmc_rank"`
	CirculatingSupply float64      `json:"circulating_supply"`
	TotalSupply       float64      `json:"total_supply"`
	MaxSupply         float64      `json:"max_supply,omitempty"`
	Tags              []*Tags      `json:"tags"`
	Platform          *Platform    `json:"platform"`
	Quote             *QuoteLatest `json:"quote"`
	LastUpdated       string       `json:"last_updated"`
	DateAdded         string       `json:"date_added"`
}

type QuoteLatest struct {
	USD *USDLatest `json:"USD"`
}

type USDLatest struct {
	Price              float64 `json:"price"`
	MarketCap          float64 `json:"market_cap"`
	MarketCapDominance float64 `json:"market_cap_dominance,omitempty"`
	PercentChange1H    float64 `json:"percent_change_1h,omitempty"`
	PercentChange24H   float64 `json:"percent_change_24h,omitempty"`
	PercentChange7D    float64 `json:"percent_change_7d,omitempty"`
	PercentChange30D   float64 `json:"percent_change_30d,omitempty"`
	Volume24H          float64 `json:"volume_24h,omitempty"`
	Volume7D           float64 `json:"volume_7d,omitempty"`
	Volume30D          float64 `json:"volume_30d,omitempty"`
	LastUpdated        string  `json:"last_updated"`
}

type ListingData struct {
	Data map[string]*Listing `json:"data"`
}

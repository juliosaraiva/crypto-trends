package model

type Crypto struct {
	ID                int                 `json:"id,omitempty"`
	Name              string              `json:"name"`
	Symbol            string              `json:"symbol"`
	Slug              string              `json:"slug"`
	CMCRank           int                 `json:"cmc_rank"`
	CirculatingSupply float64             `json:"circulating_supply"`
	TotalSupply       float64             `json:"total_supply"`
	MaxSupply         float64             `json:"max_supply,omitempty"`
	Tags              []string            `json:"tags"`
	Platform          *Platform           `json:"platform"`
	Quote             *QuoteLatest        `json:"quote"`
	LastUpdated       string              `json:"last_updated"`
	DateAdded         string              `json:"date_added"`
	Historical        []*HistoricalQuotes `json:"historical"`
}

type Cryptocurrency struct {
	ID       int       `json:"id"`
	Rank     int       `json:"rank"`
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	IsActive int       `json:"is_active"`
	Platform *Platform `json:"platform,omitempty"`
}

type CryptocurrencyData struct {
	Data []*Cryptocurrency
}

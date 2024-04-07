package types

type CategoryData struct {
	Data *Category `json:"data"`
}

type Category struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	NumTokens       int           `json:"num_tokens"`
	MarketCap       float64       `json:"market_cap"`
	MarketCapChange float64       `json:"market_cap_change"`
	Volume          float64       `json:"volume"`
	VolumeChange    float64       `json:"volume_change"`
	Coins           []*CoinLatest `json:"coins"`
}

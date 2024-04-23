package types

import "time"

type CryptocurrencyParams struct {
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Rank        int       `json:"rank"`
	MaxSupply   int       `json:"max_supply"`
	Circulating int       `json:"circulating_supply"`
	TotalSupply int       `json:"total_supply"`
	Price       float64   `json:"price"`
	TimeStamp   time.Time `json:"timestamp"`
	Trend       string    `json:"trend"`
}

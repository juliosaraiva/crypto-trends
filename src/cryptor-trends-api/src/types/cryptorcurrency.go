package types

import "time"

type CryptocurrencyParams struct {
	CoinID      int       `json:"coin_id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Rank        int       `json:"rank"`
	MaxSupply   int       `json:"max_supply"`
	Ciruclating int       `json:"circulating_supply"`
	TotalSupply int       `json:"total_supply"`
	Price       float64   `json:"price"`
	TimeStamp   time.Time `json:"timestamp"`
	Trend       string    `json:"trend"`
}

package types

import (
	"time"
)

type CryptocurrencyParams struct {
	CoinID      int       `json:"coin_id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Rank        int       `json:"rank"`
	MaxSupply   int       `json:"max_supply"`
	Ciruclating float64       `json:"circulating_supply"`
	TotalSupply int       `json:"total_supply"`
	Price       float64   `json:"price"`
	TimeStamp   time.Time `json:"timestamp"`
	Trend       string    `json:"trend"`
}

type CryptorcurrencyParamsOut struct {
	CoinID      int       `bson:"coin_id" json:"coin_id"`
	Name        string    `bson:"name" json:"name"`
	Symbol      string    `bson:"symbol" json:"symbol"`
	Rank        int       `bson:"rank" json:"rank"`
	MaxSupply   int       `bson:"max_supply" json:"max_supply"`
	Ciruclating float64       `bson:"circulating_supply" json:"circulating_supply"`
	TotalSupply int       `bson:"total_supply" json:"total_supply"`
	Price       float64   `bson:"price" json:"price"`
	TimeStamp   time.Time `bson:"timestamp" json:"timestamp"`
	Trend       string    `bson:"trend" json:"trend"`
}

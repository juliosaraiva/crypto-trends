package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cryptocurrency struct {
	ID          primitive.ObjectID `bson:"_id"`
	CoinID      int                `bson:"coin_id"`
	Name        string             `bson:"name"`
	Symbol      string             `bson:"symbol"`
	Rank        int                `bson:"rank"`
	MaxSupply   int                `bson:"max_supply"`
	Ciruclating float64                `bson:"circulating_supply"`
	TotalSupply int                `bson:"total_supply"`
	Price       float64            `bson:"price"`
	TimeStamp   time.Time          `bson:"timestamp"`
	Trend       string             `bson:"trend"`
}

func NewCryptocurrency(coinID int, name, symbol string, rank int, maxSupply int, circulating float64, totalSupply int, price float64, timestamp time.Time, trend string) (*Cryptocurrency, error) {
	c := &Cryptocurrency{
		ID:          primitive.NewObjectID(),
		CoinID:      coinID,
		Name:        name,
		Symbol:      symbol,
		Rank:        rank,
		MaxSupply:   maxSupply,
		Ciruclating: circulating,
		TotalSupply: totalSupply,
		Price:       price,
		TimeStamp:   timestamp,
		Trend:       trend,
	}

	return c, nil
}
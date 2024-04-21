package domain

import "github.com/google/uuid"

type Cryptocurrency struct {
	ID          uuid.UUID
	Name        string
	Symbol      string
	Rank        int
	MaxSupply   int
	Ciruclating int
	TotalSupply int
	Price       float64
	TimeStamp   int
	Trend       string
}

func (c *Cryptocurrency) Validate() error {
	if c.ID == "" {
		return validate.ErrorMissingField("ID")
	}
	if c.Name == "" {
		return validate.ErrorMissingField("Name")
	}
	if c.Symbol == "" {
		return validate.ErrorMissingField("Symbol")
	}
	if c.Rank == 0 {
		return validate.ErrorMissingField("Rank")
	}
	if c.MaxSupply == 0 {
		return validate.ErrorMissingField("MaxSupply")
	}
	if c.Ciruclating == 0 {
		return validate.ErrorMissingField("Ciruclating")
	}
	if c.TotalSupply == 0 {
		return validate.ErrorMissingField("TotalSupply")
	}
	if c.Price == 0 {
		return validate.ErrorMissingField("Price")
	}
	if c.TimeStamp == 0 {
		return validate.ErrorMissingField("TimeStamp")
	}
	if c.Trend == "" {
		return validate.ErrorMissingField("Trend")
	}
	return nil
}

func NewCryptocurrency(id, name, symbol string, rank, maxSupply, circulating, totalSupply, price, timestamp int, trend string) (*Cryptocurrency, error) {
	c := &Cryptocurrency{
		ID:          id,
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
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

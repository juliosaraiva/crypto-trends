package infrastructure

// TODO: Implement the function FromDBCryptocurrency
func FromDBCryptocurrency(c *Cryptocurrency) *Cryptocurrency {
	return &Cryptocurrency{
		ID:          c.ID,
		Name:        c.Name,
		Symbol:      c.Symbol,
		Rank:        c.Rank,
		MaxSupply:   c.MaxSupply,
		Ciruclating: c.Ciruclating,
		TotalSupply: c.TotalSupply,
		Price:       c.Price,
		TimeStamp:   c.TimeStamp,
		Trend:       c.Trend,
	}
}

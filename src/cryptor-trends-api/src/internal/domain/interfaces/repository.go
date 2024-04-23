package interfaces

import "github.com/juliosaraiva/crypto-trends/src/internal/domain/entities"

type ICryptocurrencyRepository interface {
	FindAll() ([]*entities.Cryptocurrency, error)
	// FindByID(coin_id int) (*entities.Cryptocurrency, error)
	Create(c *entities.Cryptocurrency) error
}

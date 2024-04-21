package interfaces

import (
	"github.com/juliosaraiva/crypto-trends/src/internal/domain/entity"
)

type ICryptocurrencyRepository interface {
	FindAll() ([]*entity.Cryptocurrency, error)
	FindByID(coin_id int) (*entity.Cryptocurrency, error)
	Create(c *entity.Cryptocurrency) error
}

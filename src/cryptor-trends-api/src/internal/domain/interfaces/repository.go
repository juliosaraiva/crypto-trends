package interfaces

import (
	"github.com/google/uuid"
)

type ICryptocurrencyRepository interface {
	FindAll() ([]*Cryptocurrency, error)
	FindByID(id uuid.UUID) (*Cryptocurrency, error)
}

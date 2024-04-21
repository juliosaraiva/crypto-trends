package infrastructure

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type GormCryptocurrency struct {
	db *gorm.DB
}

func NewCryptocurrencyRepository(db *gorm.DB) *ICryptocurrencyRepository {
	return &CryptocurrencyRepository{
		db: db,
	}
}

func (r *CryptocurrencyRepository) FindAll() ([]*Cryptocurrency, error) {
	var cryptocurrencies []*Cryptocurrency
	if err := r.db.Find(&cryptocurrencies).Error; err != nil {
		return nil, err
	}
	return cryptocurrencies, nil
}

func (r *CryptocurrencyRepository) FindByID(id uuid.UUID) (*Cryptocurrency, error) {
	var cryptocurrency Cryptocurrency
	if err := r.db.Where("id = ?", id).First(&cryptocurrency).Error; err != nil {
		return nil, err
	}
	return &cryptocurrency, nil
}

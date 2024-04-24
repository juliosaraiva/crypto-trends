package application

import (
	"context"

	"github.com/juliosaraiva/crypto-trends/src/internal/domain"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure"
	"github.com/juliosaraiva/crypto-trends/src/types"
)

type ICryptocurrencyService interface {
	FindAll(ctx context.Context) ([]*types.CryptorcurrencyParamsOut, error)
	Create(ctx context.Context, params types.CryptocurrencyParams) error
}

type CryptocurrencyService struct {
	repository infrastructure.ICryptocurrencyRepository
}

func NewCryptocurrencyService(repository infrastructure.ICryptocurrencyRepository) *CryptocurrencyService {
	return &CryptocurrencyService{
		repository: repository,
	}
}

func (c *CryptocurrencyService) FindAll(ctx context.Context) ([]*types.CryptorcurrencyParamsOut, error) {
	var cryptorCurrencyParamsOut []*types.CryptorcurrencyParamsOut
	cryptorCurrencyRepository, err := c.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, cryptorCurrency := range cryptorCurrencyRepository {
		cryptorCurrencyParamsOut = append(cryptorCurrencyParamsOut, &types.CryptorcurrencyParamsOut{
			CoinID:      cryptorCurrency.CoinID,
			Name:        cryptorCurrency.Name,
			Symbol:      cryptorCurrency.Symbol,
			Rank:        cryptorCurrency.Rank,
			MaxSupply:   cryptorCurrency.MaxSupply,
			Ciruclating: cryptorCurrency.Ciruclating,
			TotalSupply: cryptorCurrency.TotalSupply,
			Price:       cryptorCurrency.Price,
			TimeStamp:   cryptorCurrency.TimeStamp,
			Trend:       cryptorCurrency.Trend,
		})
	}

	return cryptorCurrencyParamsOut, nil
}

func (c *CryptocurrencyService) Create(ctx context.Context, params types.CryptocurrencyParams) error {
	cryptorCurrencyEntity, err := domain.NewCryptocurrency(
		params.CoinID,
		params.Name,
		params.Symbol,
		params.Rank,
		params.MaxSupply,
		params.Ciruclating,
		params.TotalSupply,
		params.Price,
		params.TimeStamp,
		params.Trend,
	)

	if err != nil {
		return err
	}

	err = c.repository.Create(ctx, cryptorCurrencyEntity)
	if err != nil {
		return err
	}

	return nil
}

package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCryptocurrency struct {
	collection *mongo.Collection
}

func NewCryptocurrencyRepository(collection *mongo.Collection) *ICryptocurrencyRepository {
	return &CryptocurrencyRepository{
		collection: collection,
	}
}

func (r *MongoCryptocurrency) FindAll(ctx context.Context) ([]*Cryptocurrency, error) {
	var cryptocurrencies []*Cryptocurrency
	filter := bson.D{{"name", "Bitcoin"}}
	cryptocurrencies, err := r.collection.Find(ctx, filter)
	if err := r.db.Find(&cryptocurrencies).Error; err != nil {
		return nil, err
	}
	return cryptocurrencies, nil
}

func (r *MongoCryptocurrency) FindByID(id uuid.UUID) (*Cryptocurrency, error) {
	var cryptocurrency Cryptocurrency
	if err := r.collection.Where("id = ?", id).First(&cryptocurrency).Error; err != nil {
		return nil, err
	}
	return &cryptocurrency, nil
}

func (r *MongoCryptocurrency) Create(ctx context.Context, cryptocurrency *Cryptocurrency) error {
	result, err := r.collection.InsertOne(ctx, cryptocurrency)
	if err != nil {
		return err
	}
	return nil
}

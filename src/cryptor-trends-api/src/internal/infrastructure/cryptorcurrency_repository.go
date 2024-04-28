package infrastructure

import (
	"context"

	"github.com/juliosaraiva/crypto-trends/src/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICryptocurrencyRepository interface {
	FindAll(ctx context.Context) ([]*domain.Cryptocurrency, error)
	Create(ctx context.Context, cryptocurrency *domain.Cryptocurrency) error
}

type MongoCryptocurrency struct {
	collection *mongo.Collection
}

func NewCryptocurrencyRepository(collection *mongo.Collection) *MongoCryptocurrency {
	return &MongoCryptocurrency{
		collection: collection,
	}
}

func (r *MongoCryptocurrency) FindAll(ctx context.Context) ([]*domain.Cryptocurrency, error) {
	var cryptocurrencies []*domain.Cryptocurrency
	filter := bson.D{{}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cryptocurrency domain.Cryptocurrency
		if err := cursor.Decode(&cryptocurrency); err != nil {
			return nil, err
		}
		cryptocurrencies = append(cryptocurrencies, &cryptocurrency)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return cryptocurrencies, nil
}

func (r *MongoCryptocurrency) Create(ctx context.Context, cryptocurrency *domain.Cryptocurrency) error {
	_, err := r.collection.InsertOne(ctx, cryptocurrency)
	if err != nil {
		return err
	}
	return nil
}

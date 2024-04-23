package repository

import (
	"context"

	"github.com/juliosaraiva/crypto-trends/src/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCryptocurrency struct {
	collection *mongo.Collection
}

func NewCryptocurrencyRepository(collection *mongo.Collection) *MongoCryptocurrency {
	return &MongoCryptocurrency{
		collection: collection,
	}
}

func (r *MongoCryptocurrency) FindAll(ctx context.Context) ([]*entities.Cryptocurrency, error) {
	var cryptocurrencies []*entities.Cryptocurrency
	filter := bson.D{{}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cryptocurrency entities.Cryptocurrency
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

// func (r *MongoCryptocurrency) FindByID(id uuid.UUID) (*Cryptocurrency, error) {
// 	var cryptocurrency Cryptocurrency
// 	if err := r.collection.Where("id = ?", id).First(&cryptocurrency).Error; err != nil {
// 		return nil, err
// 	}
// 	return &cryptocurrency, nil
// }

func (r *MongoCryptocurrency) Create(ctx context.Context, cryptocurrency *entities.Cryptocurrency) error {
	_, err := r.collection.InsertOne(ctx, cryptocurrency)
	if err != nil {
		return err
	}
	return nil
}

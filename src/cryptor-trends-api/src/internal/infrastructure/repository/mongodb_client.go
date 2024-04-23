package repository

import (
	"context"
	"fmt"

	"github.com/juliosaraiva/crypto-trends/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(config *config.Settings) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", config.MongoDBHost, config.MongoDBPort)
	ctx := context.Background()
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Collection(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func Disconnect(ctx context.Context, client *mongo.Client) error {
	return client.Disconnect(ctx)
}

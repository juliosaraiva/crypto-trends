package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/juliosaraiva/crypto-trends/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(config config.Settings) (*mongo.Client, error) {
	const uri string = fmt.Sprintf("mongodb://%s:%s/", config.MongoDBHost, config.MongoDBPort)
	return mongo.NewClient(options.Client().ApplyURI(uri))
}

func Connect(client *mongo.Client) (context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	return ctx, client.Connect(ctx)
}

func collection(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func Disconnect(ctx context.Context, client *mongo.Client) error {
	return client.Disconnect(ctx)
}

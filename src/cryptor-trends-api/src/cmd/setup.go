package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/juliosaraiva/crypto-trends/src/config"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure/repository"
)

func setup() (*config.Settings, *http.Server, error) {
	settings := config.NewSettings()

	// MongoDB connection
	mongoClient, err := repository.NewClient(settings)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	collection := repository.Collection(mongoClient, settings.MongoDBName, settings.MongoDBCollection)

	log.Printf("Connected to MongoDB on %s:%s", settings.MongoDBHost, settings.MongoDBPort)

	app = config.AppSettings{
		MongoClient:     mongoClient,
		MongoCollection: collection,
	}

	// create a new http server
	svc := &http.Server{
		Addr:              settings.WebServerPort,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	return settings, svc, nil
}

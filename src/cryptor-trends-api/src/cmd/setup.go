package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/juliosaraiva/crypto-trends/src/config"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure"
)

func setup() (*config.Settings, *http.Server, error) {
	settings := config.NewSettings()

	// MongoDB connection
	mongoClient, err := infrastructure.NewClient(settings)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	collection := infrastructure.Collection(mongoClient, settings.MongoDBName, settings.MongoDBCollection)

	log.Printf("Connected to MongoDB on %s:%s", settings.MongoDBHost, settings.MongoDBPort)

	// RabbitMQ connection
	rabbitMQClient, err := infrastructure.NewRabbitMQClient(settings)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	rabbitMQChanel, err := infrastructure.Channel(rabbitMQClient)
	if err != nil {
		log.Fatalf("Failed to create a channel: %v", err)
	}

	log.Printf("Connected to RabbitMQ on %s:%s", settings.RabbitMQHost, settings.RabbitMQPort)

	q, err := infrastructure.QueueDeclare(rabbitMQChanel, settings.RabbitMQConsumeQueueName)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// create a new config app
	app = config.AppSettings{
		MongoClient:     mongoClient,
		MongoCollection: collection,
		RabbitMQClient:  rabbitMQClient,
		RabbitMQChannel: rabbitMQChanel,
		RabbitMQQueue:   q,
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

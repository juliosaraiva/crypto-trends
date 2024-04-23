package main

import (
	"context"
	"log"
	"runtime"

	"github.com/juliosaraiva/crypto-trends/src/config"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var app config.AppSettings

func main() {
	settings, svc, err := setup()
	if err != nil {
		log.Fatalf("Error setting up the server: %v", err)
	}
	defer func(mongoClient *mongo.Client, mongoCtx context.Context) {
		if err := repository.Disconnect(mongoCtx, mongoClient); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}(app.MongoClient, context.Background())

	// print server info
	log.Printf("******************************************")
	log.Printf("** %sCryptor Trend API%s v%s built in %s", "\033[31m", "\033[0m", "v1.0.0", runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")

	// start the server
	log.Printf("Starting server on %s", settings.WebServerPort)
	if err := svc.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

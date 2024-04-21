package main

import (
	"log"
	"net/http"

	"github.com/juliosaraiva/crypto-trends/src/config"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure"
)

func setup() (*http.Server, error) {
	settings := config.NewSettings()

	// MongoDB connection
	mongoClient, err := infrastructure.NewClient(settings)
	if err != nil {
		return nil, err
	}

	mongoCtx, err := infrastructure.Connect(mongoClient)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(mongoCtx, readPref.Primary())
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to MongoDB on %s:%s", settings.MongoDBHost, settings.MongoDBPort)

	app := config.AppSettings{
		MongoClient: mongoClient,
		MongoCtx:    mongoCtx,
	}

	// create a new http server
	svc := &http.Server{
		Addr:              settings.WebServerPort,
		Handler:           nil,
		IdleTimeout:       120,
		ReadTimeout:       5,
		WriteTimeout:      10,
		ReadHeaderTimeout: 2,
	}

	return svc, nil
}

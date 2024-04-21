package main

import (
	"context"
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
	defer func(mongoClient *mongo.Client, mongoCtx context.Context) {
		if err := mongoClient.Disconnect(mongoCtx); err != nil {
			panic(err)
		}
	}(mongoClient, mongoCtx)

	err = mongoClient.Ping(mongoCtx, readPref.Primary())
	if err != nil {
		panic(err)
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

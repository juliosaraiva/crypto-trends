package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/juliosaraiva/crypto-trends/src/internal/application"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure"
	"github.com/juliosaraiva/crypto-trends/src/types"
)

func startConsumer() {
	cryptorCurrencyRepository := infrastructure.NewCryptocurrencyRepository(app.MongoCollection)
	cryptorCurrencyService := application.NewCryptocurrencyService(cryptorCurrencyRepository)

	msg, err := app.RabbitMQChannel.Consume(
		app.RabbitMQQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume message: %v", err)
	}

	var forever chan struct{}

	go func() {
		for m := range msg {
			var params types.CryptocurrencyParams
			if err := json.NewDecoder(bytes.NewReader(m.Body)).Decode(&params); err != nil {
				log.Printf("Failed to decode message: %v", err)
				continue
			}

			err := cryptorCurrencyService.Create(context.Background(), params)
			if err != nil {
				log.Printf("Failed to create cryptocurrency: %v", err)
				continue
			}
			log.Printf("Received a message: %s", m.Body)
		}
	}()
	<-forever
}

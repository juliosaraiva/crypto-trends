package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/juliosaraiva/crypto-trends/src/internal/domain/entities"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure/repository"
	"github.com/juliosaraiva/crypto-trends/src/types"
)

func routes() http.Handler {
	r := chi.NewRouter()

	// -> config middlewares
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           36000,
	}))
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/cryptocurrency", func(w http.ResponseWriter, r *http.Request) {
			cryptorCurrencyRepository := repository.NewCryptocurrencyRepository(app.MongoCollection)
			cryptoCurrent, err := cryptorCurrencyRepository.FindAll(r.Context())
			if err != nil {
				log.Printf("Failed to fetch cryptocurrencies: %v", err)
				render.Status(r, http.StatusInternalServerError)
				return
			}
			render.JSON(w, r, cryptoCurrent)
			render.Status(r, http.StatusOK)
		})
		r.Get("/cryptocurrency/{coin_id}", nil)
		r.Post("/cryptocurrency", func(w http.ResponseWriter, r *http.Request) {
			var params types.CryptocurrencyParams
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				render.Status(r, http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			// fmt.Printf("Received params: %+v\n", params)

			cryptorCurrencyEntity, err := entities.NewCryptocurrency(
				params.CoinID,
				params.Name,
				params.Symbol,
				params.Rank,
				params.MaxSupply,
				params.Ciruclating,
				params.TotalSupply,
				params.Price,
				params.TimeStamp,
				params.Trend,
			)

			if err != nil {
				log.Printf("Failed to create cryptocurrency entity: %v", err)
				render.Status(r, http.StatusInternalServerError)
				return
			}

			fmt.Printf("Received entity: %+v\n", cryptorCurrencyEntity)

			cryptorCurrencyRepository := repository.NewCryptocurrencyRepository(app.MongoCollection)
			if err := cryptorCurrencyRepository.Create(r.Context(), cryptorCurrencyEntity); err != nil {
				log.Printf("Failed to create cryptocurrency: %v", err)
				render.Status(r, http.StatusInternalServerError)
				return
			}

			render.Status(r, http.StatusCreated)
		})

		// -> healthcheck
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
		})
	})

	log.Printf("Routes configured")

	return r
}

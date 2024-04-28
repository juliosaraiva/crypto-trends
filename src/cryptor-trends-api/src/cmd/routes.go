package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/juliosaraiva/crypto-trends/src/internal/application"
	"github.com/juliosaraiva/crypto-trends/src/internal/infrastructure"
)

func routes() http.Handler {
	// -> setup handler
	cryptorCurrencyRepository := infrastructure.NewCryptocurrencyRepository(app.MongoCollection)
	cryptorCurrencyService := application.NewCryptocurrencyService(cryptorCurrencyRepository)
	cryptoCurrencyHandler := application.NewCryptocurrencyHandler(cryptorCurrencyService)

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
		// -> cryptocurrency
		r.Route("/cryptocurrency", func(r chi.Router) {
			r.Get("/", cryptoCurrencyHandler.FindAll)
			r.Post("/", cryptoCurrencyHandler.Create)
		})
		// -> healthcheck
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
		})
	})

	log.Printf("Routes configured")

	return r
}

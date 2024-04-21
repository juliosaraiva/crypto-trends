package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func routes() {
	r := chi.NewRouter()

	r.Get("/cryptocurrency", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})
	r.Get("/cryptocurrency/{coin_id}", nil)
	r.Post("/cryptocurrency", func(w http.ResponseWriter, r *http.Request) {

		render.Status(r, http.StatusCreated)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})
}

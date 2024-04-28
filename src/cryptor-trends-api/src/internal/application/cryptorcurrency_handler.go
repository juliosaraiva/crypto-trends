package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/juliosaraiva/crypto-trends/src/types"
)

type CryptocurrencyHandler struct {
	service ICryptocurrencyService
}

func NewCryptocurrencyHandler(service ICryptocurrencyService) *CryptocurrencyHandler {
	return &CryptocurrencyHandler{
		service: service,
	}
}

func (c *CryptocurrencyHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	cryptoCurrent, err := c.service.FindAll(r.Context())
	if err != nil {
		log.Printf("Failed to find cryptocurrencies: %v", err)
		render.Status(r, http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, cryptoCurrent)
	render.Status(r, http.StatusOK)
}

func (c *CryptocurrencyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var params types.CryptocurrencyParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := c.service.Create(r.Context(), params)
	if err != nil {
		log.Printf("Failed to create cryptocurrency: %v", err)
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
}

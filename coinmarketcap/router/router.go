package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/historical", handler.GetCryptoHistoricalPrices)
	api.Get("/crypto", handler.GetCrypto)
	api.Get("/map", handler.ListCryptocurrencies)
}
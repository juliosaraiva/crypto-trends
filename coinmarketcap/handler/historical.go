package handler

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
)

type HistoricalQueryParam struct {
	ID          string
	Symbol      string
	TimeStart   string
	TimeEnd     string
	Count       int
	Interval    string
	Aux         string
	SkipInvalid bool
}

func GetCryptoHistoricalPrices(c *fiber.Ctx) error {
	queryParams := url.Values{}
	queries := c.Queries()

	for k, v := range queries {
		queryParams[k] = []string{v}
	}

	reqHeaders := c.GetReqHeaders()

	historicalPrices, err := coinmarketcap.GetHistoricalPrices(c.Context(), queryParams, reqHeaders)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.JSON(historicalPrices)
}

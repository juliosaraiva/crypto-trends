package handler

import (
	"encoding/json"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
	"github.com/juliosaraiva/crypto-trends/types"
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
		var jsonError types.JSONError
		json.Unmarshal([]byte(err.Error()), &jsonError)
		return c.Status(fiber.StatusBadRequest).JSON(jsonError)
	}

	return c.JSON(historicalPrices)
}

func GetOHLCVHistoricalPrices(c *fiber.Ctx) error {
	queryParams := url.Values{}
	queries := c.Queries()

	for k, v := range queries {
		queryParams[k] = []string{v}
	}

	reqHeaders := c.GetReqHeaders()

	OHLCVHistoricalPrices, err := coinmarketcap.GetOHLCVHistoricalPrices(c.Context(), queryParams, reqHeaders)
	if err != nil {
		var jsonError types.JSONError
		json.Unmarshal([]byte(err.Error()), &jsonError)
		return c.Status(fiber.StatusBadRequest).JSON(jsonError)
	}

	return c.JSON(OHLCVHistoricalPrices)
}

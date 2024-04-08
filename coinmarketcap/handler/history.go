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

func GetHistory(c *fiber.Ctx) error {
	q := url.Values{}
	queries := c.Queries()

	for k, v := range queries {
		q[k] = []string{v}
	}

	history, err := coinmarketcap.GetHistory(c.Context(), &q)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.JSON(&history)
}

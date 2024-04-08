package handler

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
)

func GetListing(c *fiber.Ctx) error {
	var q url.Values = url.Values{}
	queries := c.Queries()
	for k, v := range queries {
		q[k] = []string{v}
	}
	filtered, unfiltered, err := coinmarketcap.GetLatest(c.Context(), q)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err})
	} else if len(filtered) > 0 {
		return c.JSON(filtered)
	} else {
		return c.JSON(unfiltered)
	}
}

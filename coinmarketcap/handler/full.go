package handler

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
)

func GetAllCryptos(c *fiber.Ctx) error {
	var q url.Values = url.Values{}

	filtered, unfiltered, err := coinmarketcap.GetLatest(c.Context(), q)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	} else if len(filtered) > 0 {
		return c.JSON(filtered)
	} else {
		return c.JSON(unfiltered)
	}
}

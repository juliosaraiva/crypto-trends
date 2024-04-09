package handler

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
)

func GetAllCryptos(c *fiber.Ctx) error {
	var q url.Values = url.Values{}
	queries := c.Queries()

	for k, v := range queries {
		q[k] = []string{v}
	}

	reqHeaders := c.GetReqHeaders()

	cryptos, err := coinmarketcap.GetCrypto(c.Context(), q, reqHeaders)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.JSON(cryptos)
}

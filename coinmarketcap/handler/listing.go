package handler

import (
	"encoding/json"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
	"github.com/juliosaraiva/crypto-trends/types"
)

func GetListing(c *fiber.Ctx) error {
	var q url.Values = url.Values{}
	queries := c.Queries()

	for k, v := range queries {
		q[k] = []string{v}
	}

	reqHeaders := c.GetReqHeaders()
	tokens, err := coinmarketcap.GetListing(c.Context(), q, reqHeaders)
	if err != nil {
		var jsonError types.JSONError
		json.Unmarshal([]byte(err.Error()), &jsonError)
		return c.Status(fiber.StatusBadRequest).JSON(jsonError)
	}

	return c.JSON(tokens)
}

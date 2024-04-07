package handler

import (
	"github.com/gofiber/fiber/v2"
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
	return c.JSON(fiber.Map{"msg": "OK"})
}

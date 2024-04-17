package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap/router"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{"status_code": code, "error": err.Error()})
	},
}

func main() {
	var port string = ":8000"
	app := fiber.New(config)
	app.Use(cors.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}

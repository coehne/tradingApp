package main

import (
	"log"

	"github.com/dakicka/tradingApp/api/db"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the TradingApp API")
}

func main() {
	db.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}

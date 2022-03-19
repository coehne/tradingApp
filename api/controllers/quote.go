package controllers

import (
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/gofiber/fiber/v2"
)

func GetQuote(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// Get stock info from iex cloud API
	quote, err := integration.GetStockInfo(data["symbol"])
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Could not get stock data",
		})
	}

	return c.JSON(quote)
}

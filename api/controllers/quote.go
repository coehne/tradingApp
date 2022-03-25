package controllers

import (
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/gofiber/fiber/v2"
)

func GetQuote(c *fiber.Ctx) error {
	symbol := c.Params("symbol")

	// Get stock info from iex cloud API
	quote, err := integration.GetStockInfo(symbol)
	if err != nil {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "invalid symbol or stock data",
		})
	}
	return c.JSON(quote)
}

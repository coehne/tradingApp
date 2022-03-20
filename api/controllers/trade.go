package controllers

import (
	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/gofiber/fiber/v2"
)

type CreateTradeRequest struct {
	Symbol string `json:"symbol"`
	Qty    int    `json:"qty"`
}

func CreateTrade(c *fiber.Ctx) error {
	// Get user from token
	user, err := GetUserFromToken(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get data from request
	var data CreateTradeRequest
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Get current stock price
	stock, err := integration.GetStockInfo(data.Symbol)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// Build trade
	var trade models.Trade
	trade.Price = stock.LatestPrice
	trade.UserID = user.ID
	trade.Qty = data.Qty
	trade.Symbol = data.Symbol

	// Check if user has enough cash
	// Get cash
	var cash float64
	transactions := []models.Transaction{}
	database.DB.Find(&transactions, "user_id = ?", user.ID)
	for _, tx := range transactions {
		cash += tx.Amount
	}

	// Build transaction
	var transaction models.Transaction
	transaction.UserID = user.ID
	transaction.Amount = trade.Price * float64(trade.Qty) * (-1)

	// Return 400 if not enough cash for transaction
	if trade.Qty > 0 && cash < transaction.Amount {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "not enough cash",
		})
	}

	// Return 400 if not enough stocks to sell
	if transaction.Amount > 0 {
		var stocks int
		trades := []models.Trade{}
		database.DB.Find(&trades, "user_id = ?", user.ID)
		for _, trade := range trades {
			if data.Symbol == trade.Symbol {
				stocks += trade.Qty

			}
		}
		if stocks < 0 {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "not enough stocks",
			})
		}

	}

	// Insert into db
	database.DB.Create(&trade)
	database.DB.Create(&transaction)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetTrades(c *fiber.Ctx) error {
	// Get user from token
	user, err := GetUserFromToken(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	trades := []models.Trade{}
	database.DB.Find(&trades, "user_id = ?", user.ID)

	return c.Status(fiber.StatusOK).JSON(trades)

}

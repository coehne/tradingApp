package controllers

import (
	"errors"

	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/integration"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/gofiber/fiber/v2"
)

type CreateTradeRequest struct {
	Symbol string `json:"symbol"`
	Qty    int    `json:"qty" validate:"required"`
}

func FindTrade(id int, trade *models.Trade) error {
	database.DB.Find(&trade, "id = ?", id)
	if trade.ID == 0 {
		return errors.New("trade does not exist")
	}
	return nil
}

func CreateTrade(c *fiber.Ctx) error {
	// Get user from token
	user, err := GetUserFromToken(c)
	// Return 401 if invalid or no token provided
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
			"message": "internal server error",
		})
	}
	// Build trade
	var trade models.Trade
	trade.Price = stock.LatestPrice
	trade.UserID = user.ID
	trade.Qty = data.Qty
	trade.Symbol = data.Symbol
	trade.CompanyName = stock.CompanyName

	// Build transaction
	var transaction models.Transaction
	transaction.UserID = user.ID
	transaction.Amount = trade.Price * float64(trade.Qty) * (-1)

	// Check if user has enough cash
	// Get cash
	var cash float64
	transactions := []models.Transaction{}
	database.DB.Find(&transactions, "user_id = ?", user.ID)
	for _, tx := range transactions {
		cash += tx.Amount
	}

	// Return 400 if not enough cash for transaction
	if trade.Qty > 0 && cash < transaction.Amount {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "not enough cash",
		})
	}

	// Return 400 if not enough stocks to sell
	if transaction.Amount > 0 {
		// Check how many stonks of that symbol are in the depot
		var stocks int
		trades := []models.Trade{}
		database.DB.Find(&trades, "user_id = ?", user.ID)
		for _, trade := range trades {
			if data.Symbol == trade.Symbol {
				stocks += trade.Qty

			}
		}
		// If the amount of stocks - the sell order is negativ, return 400
		if stocks+trade.Qty < 0 {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "not enough stocks",
			})
		}

	}

	// Insert trade into db
	database.DB.Create(&trade)

	// Adjust add generated trade ID as FK to transaction
	transaction.TradeID = trade.ID

	// Insert transaction into db
	database.DB.Create(&transaction)

	// Return 200
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetTrades(c *fiber.Ctx) error {
	// Get user from token
	user, err := GetUserFromToken(c)
	// Return 401 if invalid or no token provided
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
func GetTrade(c *fiber.Ctx) error {
	// Get user from token
	user, err := GetUserFromToken(c)
	// Return 401 if invalid or no token provided
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	// Get tradeId from params
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	trade := models.Trade{}
	if err := FindTrade(id, &trade); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.DB.Find(&trade, "user_id = ?", user.ID)

	return c.Status(fiber.StatusOK).JSON(trade)

}

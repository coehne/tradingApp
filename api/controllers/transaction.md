/* package controllers

import (
	"time"

	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/gofiber/fiber/v2"
)

type CreateTransactionRequest struct {
	Amount float64 `json:"amount"`
}

type GetTransactionsResponse struct {
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

func CreateTransaction(c *fiber.Ctx) error {

	// Get user from token
	user, err := GetUserFromToken(c)
	// Return 401 if invalid or no token provided
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get amount from request
	var data CreateTransactionRequest
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var transaction models.Transaction

	transaction.UserID = user.ID
	transaction.Amount = data.Amount
	database.DB.Create(&transaction)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func GetTransactions(c *fiber.Ctx) error {
	// Get user from token
	user, err := GetUserFromToken(c)
	// Return 401 if invalid or no token provided
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	transactions := []models.Transaction{}
	database.DB.Find(&transactions, "user_id = ?", user.ID)

	// Return 200
	return c.Status(fiber.StatusOK).JSON(transactions)

}
*/
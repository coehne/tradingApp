package controllers

import (
	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/models"
	"github.com/gofiber/fiber/v2"
)

func DepositCash(c *fiber.Ctx) error {

	// Get user from token
	user, err := GetUserFromToken(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get amount from request
	var data map[string]int64
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var transaction models.Transaction

	transaction.UserID = user.ID
	transaction.Value = data["value"]
	database.DB.Create(&transaction)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

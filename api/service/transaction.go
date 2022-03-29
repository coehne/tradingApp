package service

import (
	"github.com/dakicka/tradingApp/api/auth"
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
)

func (s Service) CreateTransaction(ctx *fiber.Ctx, amount float64) (entity.Transaction, error) {
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return entity.Transaction{}, fiber.NewError(fiber.ErrInternalServerError.Code, "could not get user id from token")
	}

	if amount <= 0 {
		// Check if user has enough cash to withdraw
		cash, err := s.GetCashForUserId(userId)
		if err != nil {
			return entity.Transaction{}, fiber.NewError(fiber.StatusInternalServerError, "could not get cash for user")
		}
		if cash < (amount * (-1)) {
			return entity.Transaction{}, fiber.NewError(fiber.StatusBadRequest, "not enough cash for transaction")
		}

	}

	// Build transaction
	tx := entity.Transaction{
		Amount: amount,
		UserID: userId,
	}

	// Send transaction to repo
	tx, err = s.transactions.CreateTransaction(tx)
	if err != nil {
		return entity.Transaction{}, err
	}

	return tx, nil
}

func (s Service) GetAllForUserId(ctx *fiber.Ctx) ([]entity.Transaction, error) {

	// Get user id from context/cookie/token
	userId, err := auth.GetUserIdFromToken(ctx)
	if err != nil {
		return []entity.Transaction{}, fiber.NewError(fiber.StatusInternalServerError, "could not get user id from token")
	}
	// Get transactions from repository
	transactions, err := s.transactions.GetAllForUserId(userId)
	if err != nil {
		return []entity.Transaction{}, err
	}
	return transactions, nil
}

func (s Service) GetCashForUserId(userId uint) (float64, error) {
	var cash float64
	txns, err := s.transactions.GetAllForUserId(userId)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "could not get query from repo")
	}
	for _, tx := range txns {
		cash += tx.Amount
	}
	return cash, nil
}

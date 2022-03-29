package usecase

import (
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
)

type user interface {
	RegisterUser(firstName, email, password string) (entity.User, error)
	GetUserFromId(id uint) (entity.User, error)
	Login(email, password string) (entity.User, error)
}

type trade interface {
	CreateTrade(ctx *fiber.Ctx, qty int, symbol string) (entity.Trade, error)
	GetTradesByUserId(ctx *fiber.Ctx) ([]entity.Trade, error)
	GetTradeById(ctx *fiber.Ctx) (entity.Trade, error)
	GetTradesForDepot(ctx *fiber.Ctx) ([]entity.Trade, error)
}

type transaction interface {
	CreateTransaction(ctx *fiber.Ctx, amount float64) (entity.Transaction, error)
	GetAllForUserId(ctx *fiber.Ctx) ([]entity.Transaction, error)
}

type UseCases interface {
	user
	trade
	transaction
}

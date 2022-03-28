package usecase

import (
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
)

type user interface {
	RegisterUser(ctx fiber.Ctx, firstName, email, password string) (entity.User, error)
	GetUserFromId(ctx fiber.Ctx, id uint) (entity.User, error)
	Login(ctx fiber.Ctx, email, password string) (entity.User, error)
	Logout(ctx fiber.Ctx) error
}

type UseCases interface {
	user
}

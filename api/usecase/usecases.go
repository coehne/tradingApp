package usecase

import (
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
)

type user interface {
	RegisterUser(firstName, email, password string) (entity.User, error)
	GetUserFromId(id uint) (entity.User, error)
	Login(email, password string) (entity.User, error)
	Logout(ctx *fiber.Ctx) error
}

type UseCases interface {
	user
}

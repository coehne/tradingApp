package repository

import (
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/gofiber/fiber/v2"
)

type Users interface {
	Create(ctx fiber.Ctx, user entity.User) (entity.User, error)
	Get(ctx fiber.Ctx, user entity.User) (entity.User, error)
	GetByEmail(ctx fiber.Ctx, user entity.User) (entity.User, error)
}

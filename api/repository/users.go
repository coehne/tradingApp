package repository

import (
	"github.com/dakicka/tradingApp/api/entity"
)

type Users interface {
	Create(user entity.User) (entity.User, error)
	Get(user entity.User) (entity.User, error)
	GetByEmail(user entity.User) (entity.User, error)
}

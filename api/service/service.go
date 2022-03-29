package service

import (
	"github.com/dakicka/tradingApp/api/db"
	"github.com/dakicka/tradingApp/api/repository"
	"github.com/dakicka/tradingApp/api/usecase"
)

// Service wraps the required repositories and implements the defined use cases.
type Service struct {
	users repository.Users
}

// New initializes and returns a new service with the given repositories.
func New(
	ur repository.Users) usecase.UseCases {
	return &Service{ur}
}

func Init(db *db.GormDB) usecase.UseCases {
	return Service{
		users: repository.NewUsersSQLRepo(db),
	}
}

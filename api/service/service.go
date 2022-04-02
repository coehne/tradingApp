package service

import (
	"github.com/dakicka/tradingApp/api/package/db"
	"github.com/dakicka/tradingApp/api/repository"
	"github.com/dakicka/tradingApp/api/usecase"
)

// Service wraps the required repositories and implements the defined use cases.
type Service struct {
	users        repository.Users
	trades       repository.Trades
	transactions repository.Transactions
}

// New initializes and returns a new service with the given repositories.
func New(
	ur repository.Users,
	tr repository.Trades,
	tx repository.Transactions) usecase.UseCases {
	return &Service{ur, tr, tx}
}

func Init(db *db.GormDB) usecase.UseCases {
	return Service{
		users:        repository.NewUsersSQLRepo(db),
		trades:       repository.NewTradesSQLRepo(db),
		transactions: repository.NewTransactionSQLRepo(db),
	}
}

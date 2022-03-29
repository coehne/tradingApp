package repository

import "github.com/dakicka/tradingApp/api/entity"

type Transactions interface {
	CreateTransaction(tx entity.Transaction) (entity.Transaction, error)
	GetAllForUserId(userId uint) ([]entity.Transaction, error)
}

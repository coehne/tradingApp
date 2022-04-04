package repository

import "github.com/dakicka/tradingApp/api/entity"

type Transactions interface {
	Create(tx entity.Transaction) (entity.Transaction, error)
	GetAllByUserId(userId uint) ([]entity.Transaction, error)
}

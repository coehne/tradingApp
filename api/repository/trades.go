package repository

import "github.com/dakicka/tradingApp/api/entity"

type Trades interface {
	Create(trade entity.Trade) (entity.Trade, error)
	GetAllByUserId(tuserId uint) ([]entity.Trade, error)
}

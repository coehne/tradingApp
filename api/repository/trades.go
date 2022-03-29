package repository

import "github.com/dakicka/tradingApp/api/entity"

type Trades interface {
	Create(trade entity.Trade) (entity.Trade, error)
	GetAllByUserId(userId uint) ([]entity.Trade, error)
	GetById(userId uint, tradeId uint) (entity.Trade, error)
	GetDepot(userId uint) ([]entity.Trade, error)
}

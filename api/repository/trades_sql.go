package repository

import (
	"github.com/dakicka/tradingApp/api/db"
	"github.com/dakicka/tradingApp/api/entity"
)

// UsersSQL wraps the SQL DB and implements the required operations.
type TradesSQL struct {
	db.GormDB
}

// NewTradesSQLRepo instanciates and returns a new trades repository.
func NewTradesSQLRepo(db *db.GormDB) Trades {
	return TradesSQL{*db}
}

func (r TradesSQL) Create(trade entity.Trade) (entity.Trade, error) {
	// Insert into DB
	result := r.DB.Create(&trade)

	// Check for errors during insertion
	if result.Error != nil {
		return entity.Trade{}, result.Error
	}

	return trade, nil
}

func (r TradesSQL) GetAllByUserId(userId uint) ([]entity.Trade, error) {

	trades := []entity.Trade{}

	result := r.DB.Find(&trades, "user_id = ?", userId)
	// Check for errors during insertion
	if result.Error != nil {
		return []entity.Trade{}, result.Error
	}

	return trades, nil
}

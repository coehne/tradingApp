package repository

import (
	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/package/db"
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

	// Check for errors during query
	if result.Error != nil {
		return []entity.Trade{}, result.Error
	}

	return trades, nil
}

func (r TradesSQL) GetById(userId uint, tradeId uint) (entity.Trade, error) {

	// Use the userId from the cookie to make sure you can only query your own trades
	trade := entity.Trade{}

	result := r.DB.First(&trade, "user_id = ? AND id = ?", userId, tradeId)

	// Check for errors during query
	if result.Error != nil {
		return entity.Trade{}, result.Error
	}

	return trade, nil

}

func (r TradesSQL) GetDepotByUserId(userId uint) ([]entity.Trade, error) {

	var trades []entity.Trade

	result := r.DB.Model(&entity.Trade{}).Select("company_name, symbol, sum(qty) as qty").Group("symbol, company_name").Where("user_id = ?", userId).Find(&trades)

	// Check for errors during query
	if result.Error != nil {
		return []entity.Trade{}, result.Error
	}
	return trades, nil
}

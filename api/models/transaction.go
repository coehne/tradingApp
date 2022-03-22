package models

import "time"

type Transaction struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uint      `json:"userId"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	Amount    float64   `json:"amount"`
	TradeID   uint      `json:"-"`
	Trade     Trade     `json:"Trade" gorm:"foreignKey:TradeID"`
}

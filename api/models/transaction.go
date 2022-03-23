package models

import "time"

type Transaction struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uint      `json:"userId"`
	User      User      `json:"-"`
	Amount    float64   `json:"amount"`
	TradeID   uint      `json:"tradeId"`
	Trade     Trade     `json:"-"`
}

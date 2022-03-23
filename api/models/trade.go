package models

import "time"

type Trade struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	Symbol    string    `json:"symbol"`
	Qty       int       `json:"qty"`
	Price     float64   `json:"price"`
}

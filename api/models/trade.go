package models

import "time"

type Trade struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"-"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	Symbol    string    `json:"symbol"`
	Qty       int       `json:"qty"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

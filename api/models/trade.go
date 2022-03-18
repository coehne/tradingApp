package models

import "time"

type Trade struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserId    int     `json:"userId"`
	User      User    `gorm:"foreignKey:UserId"`
	Symbol    string  `json:"symbol"`
	Qty       int     `json:"qty"`
	Price     float32 `json:"price"`
	CreatedAt time.Time
}

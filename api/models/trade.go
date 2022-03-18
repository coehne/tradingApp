package models

import "time"

type Trade struct {
	ID        uint `gorm:"primaryKey"`
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	Symbol    string
	Qty       int
	Price     float32
	CreatedAt time.Time
}

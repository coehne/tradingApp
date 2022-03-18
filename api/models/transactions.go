package models

import "time"

type Transaction struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserId    int     `json:"userId"`
	User      User    `gorm:"foreignKey:UserId"`
	Value     float32 `json:"value"`
	CreatedAt time.Time
}

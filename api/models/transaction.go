package models

import "time"

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userId"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	Value     float32   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

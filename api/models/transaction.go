package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID  uint    `json:"userId"`
	User    User    `json:"-" gorm:"foreignKey:UserID"`
	Amount  float64 `json:"amount"`
	TradeID uint    `json:"-"`
	Trade   Trade   `json:"Trade" gorm:"foreignKey:TradeID"`
}

package models

import (
	"gorm.io/gorm"
)

type Trade struct {
	gorm.Model
	UserID uint    `json:"-"`
	User   User    `json:"-" gorm:"foreignKey:UserID"`
	Symbol string  `json:"symbol"`
	Qty    int     `json:"qty"`
	Price  float64 `json:"price"`
}

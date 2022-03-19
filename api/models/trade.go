package models

import (
	"gorm.io/gorm"
)

type Trade struct {
	gorm.Model
	UserID int     `json:"-"`
	User   User    `json:"-" gorm:"foreignKey:UserID"`
	Symbol string  `json:"symbol"`
	Qty    int     `json:"qty"`
	Price  float32 `json:"price"`
}

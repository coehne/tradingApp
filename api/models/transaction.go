package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID int     `json:"userId"`
	User   User    `json:"-" gorm:"foreignKey:UserID"`
	Value  float32 `json:"value"`
}

package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID uint  `json:"userId"`
	User   User  `json:"-" gorm:"foreignKey:UserID"`
	Value  int64 `json:"value"`
}

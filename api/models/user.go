package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Hash      []byte `json:"-"`
	FirstName string `json:"firstName"`
}

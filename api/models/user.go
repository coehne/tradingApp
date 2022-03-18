package models

import "time"

type User struct {
	ID        uint
	Email     string `gorm:"unique"`
	Hash      []byte
	FirstName string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

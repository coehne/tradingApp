package models

import "time"

type User struct {
	ID        uint       `json:"id"`
	Email     string     `gorm:"unique"`
	Hash      []byte     `json:"-"`
	FirstName string     `json:"firstName"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

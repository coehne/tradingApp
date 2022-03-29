package entity

import (
	"time"
)

type User struct {
	ID        uint       `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time  `json:"-" gorm:"not null"`
	UpdatedAt *time.Time `json:"-" gorm:"not null"`
	Email     string     `json:"email,omitempty" gorm:"unique;not null"`
	Hash      []byte     `json:"-" gorm:"not null"`
	FirstName string     `json:"firstName,omitempty" gorm:"not null"`

	// TODO: Make virtual fields possible with auto calc of values
}

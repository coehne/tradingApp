package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateddAt"`
	Email     string    `gorm:"unique"`
	Hash      []byte    `json:"-"`
	FirstName string    `json:"firstName"`

	// TODO: Make virtual fields possible with auto calc of values
}

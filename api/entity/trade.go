package entity

import "time"

type Trade struct {
	ID          uint      `json:"id,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UserID      uint      `json:"-"`
	User        User      `json:"-"`
	Symbol      string    `json:"symbol"`
	CompanyName string    `json:"companyName"`
	Qty         int       `json:"qty"`
	Price       float64   `json:"price"`
}

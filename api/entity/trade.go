package entity

import "time"

type Trade struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UserID      uint      `json:"-"`
	User        User      `json:"-"`
	Symbol      string    `json:"symbol"`
	CompanyName string    `json:"companyName"`
	Qty         int       `json:"qty"`
	Price       float64   `json:"price"`
}

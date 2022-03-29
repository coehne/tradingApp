package api

type meRes struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
}

type depotRes struct {
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"companyName"`
	Qty         int64   `json:"qty"`
	Price       float64 `json:"price"`
}

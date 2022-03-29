package api

type registerReq struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createTradeReq struct {
	Symbol string `json:"symbol"`
	Qty    int    `json:"qty"`
}

type createTransactionReq struct {
	Amount float64 `json:"amount"`
}

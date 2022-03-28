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

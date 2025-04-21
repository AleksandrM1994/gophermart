package api

type GetBalanceResponse struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}

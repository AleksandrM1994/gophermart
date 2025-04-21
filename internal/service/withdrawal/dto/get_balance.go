package dto

type GetBalanceRequest struct {
	UserID string
}

type GetBalanceResponse struct {
	Current   int
	Withdrawn int
}

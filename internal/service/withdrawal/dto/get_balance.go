package dto

type GetBalanceRequest struct {
	UserID string
}

type GetBalanceResponse struct {
	Current   float32
	Withdrawn float32
}

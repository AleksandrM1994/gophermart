package dto

type MakeWithdrawalRequest struct {
	Order  string
	Sum    float32
	UserID string
}

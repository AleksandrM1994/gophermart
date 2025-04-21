package dto

type MakeWithdrawalRequest struct {
	Order  string
	Sum    int
	UserID string
}

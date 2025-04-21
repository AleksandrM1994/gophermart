package dto

import "time"

type GetWithdrawalsInfoRequest struct {
	UserID string
}

type GetWithdrawalsInfoResponse struct {
	withdrawals []*Withdrawal
}

type Withdrawal struct {
	Order       string
	Sum         int
	ProcessedAt time.Time
}

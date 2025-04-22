package api

import "time"

type GetWithdrawalsInfoResponse struct {
	Withdrawals []*Withdrawal
}

type Withdrawal struct {
	Order       string    `json:"order"`
	Sum         float32   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

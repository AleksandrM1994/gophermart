package api

import "time"

type GetWithdrawalsInfoResponse struct {
	OrderID     string    `json:"order"`
	Sum         float32   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

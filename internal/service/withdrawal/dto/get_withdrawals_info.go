package dto

import "time"

type GetWithdrawalsInfoRequest struct {
	UserID string
}

type GetWithdrawalsInfoResponse struct {
	OrderID     string
	Sum         float32
	ProcessedAt time.Time
}

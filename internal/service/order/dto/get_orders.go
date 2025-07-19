package dto

import "time"

type GetOrdersRequest struct {
	UserID string
}

type GetOrdersResponse struct {
	Number     string
	Status     string
	Accrual    float32
	UploadedAt time.Time
}

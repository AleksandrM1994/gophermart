package dto

import "time"

type GetOrdersRequest struct {
	UserID string
}

type GetOrdersResponse struct {
	Orders []*Order
}

type Order struct {
	Number     string
	Status     string
	Accrual    int
	UploadedAt time.Time
}

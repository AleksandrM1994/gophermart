package dto

type GetOrderInfoRequest struct {
	Order string `json:"order"`
}

type GetOrderInfoResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual,omitempty"`
}

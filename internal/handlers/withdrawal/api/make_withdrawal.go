package api

type MakeWithdrawalRequest struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

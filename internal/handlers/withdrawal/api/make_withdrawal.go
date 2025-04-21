package api

type MakeWithdrawalRequest struct {
	Order string `json:"order"`
	Sum   int    `json:"sum"`
}

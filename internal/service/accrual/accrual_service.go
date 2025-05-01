package accrual

import (
	"context"

	"github.com/gophermart/internal/service/accrual/dto"
)

type AccrualService interface {
	GetOrderInfo(ctx context.Context, req *dto.GetOrderInfoRequest) (*dto.GetOrderInfoResponse, error)
}

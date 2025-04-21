package withdrawal

import (
	"context"

	"github.com/gophermart/internal/service/withdrawal/dto"
)

type WithdrawalService interface {
	GetBalance(ctx context.Context, req *dto.GetBalanceRequest) (*dto.GetBalanceResponse, error)
	MakeWithdrawal(ctx context.Context, req *dto.MakeWithdrawalRequest) error
	GetWithdrawalsInfo(ctx context.Context, req *dto.GetWithdrawalsInfoRequest) (*dto.GetWithdrawalsInfoResponse, error)
}

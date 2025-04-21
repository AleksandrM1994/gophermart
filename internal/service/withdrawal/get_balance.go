package withdrawal

import (
	"context"
	"fmt"

	"github.com/gophermart/internal/service/withdrawal/dto"
)

func (s *WithdrawalServiceImpl) GetBalance(ctx context.Context, req *dto.GetBalanceRequest) (*dto.GetBalanceResponse, error) {
	user, errGetUserByID := s.userRepo.GetUserByID(ctx, req.UserID)
	if errGetUserByID != nil {
		return nil, fmt.Errorf("userRepo.GetUserByID: %w", errGetUserByID)
	}

	withdrawals, errGetWithdrawalByUserID := s.withdrawalRepo.GetWithdrawalByUserID(ctx, req.UserID)
	if errGetWithdrawalByUserID != nil {
		return nil, fmt.Errorf("withdrawalRepo.GetWithdrawalByUserID: %w", errGetWithdrawalByUserID)
	}

	var withdrawalSum int
	for _, withdrawal := range withdrawals {
		withdrawalSum += withdrawal.Sum
	}

	return &dto.GetBalanceResponse{
		Current:   user.Balance,
		Withdrawn: withdrawalSum,
	}, nil
}

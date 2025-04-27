package withdrawal

import (
	"context"
	"fmt"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/withdrawal/dto"
)

func (s *WithdrawalServiceImpl) MakeWithdrawal(ctx context.Context, req *dto.MakeWithdrawalRequest) error {
	s.lg.Infow("MAKE WITHDRAWALS REQUEST", "make_withdrawals_request", req)

	if !service.LunaCheck(req.Order) {
		return fmt.Errorf("LunaCheck for order: %w", custom_errs.ErrWrongFormat)
	}

	user, errGetUserByID := s.userRepo.GetUserByID(ctx, req.UserID)
	if errGetUserByID != nil {
		return fmt.Errorf("userRepo.GetUserByID:%w", errGetUserByID)
	}

	if user.Balance < req.Sum {
		return custom_errs.ErrNotFunds
	}

	err := s.withdrawalRepo.CreateWithdrawal(ctx, &repository.Withdrawal{
		OrderID:     req.Order,
		Sum:         req.Sum,
		ProcessedAt: service.DatePtr(time.Now()),
		UserID:      req.UserID,
	})
	if err != nil {
		return fmt.Errorf("withdrawalRepo.CreateWithdrawal: %w", err)
	}

	errUpdateUserByID := s.userRepo.UpdateUserByID(ctx, req.UserID, func(currentUser *repository.User) error {
		currentUser.Balance -= req.Sum
		return nil
	})
	if errUpdateUserByID != nil {
		return fmt.Errorf("userRepo.UpdateUserByID: %w", errUpdateUserByID)
	}

	return nil
}

package withdrawal

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/gophermart/internal/service/withdrawal/dto"
)

func (s *WithdrawalServiceImpl) GetWithdrawalsInfo(ctx context.Context, req *dto.GetWithdrawalsInfoRequest) ([]*dto.GetWithdrawalsInfoResponse, error) {
	s.lg.Infow("GET WITHDRAWALS INFO REQUEST", "get_withdrawals_info_request", req)

	withdrawal, err := s.withdrawalRepo.GetWithdrawalByUserID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("withdrawalRepo.GetWithdrawalByUserID: %w", err)
	}

	res := make([]*dto.GetWithdrawalsInfoResponse, 0)
	errCopy := copier.Copy(&res, withdrawal)
	if errCopy != nil {
		return nil, fmt.Errorf("copier.Copy: %w", errCopy)
	}

	return res, nil
}

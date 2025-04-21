package order

import (
	"context"
	"fmt"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/order/dto"
)

func (s *OrderServiceImpl) GetOrders(ctx context.Context, req *dto.GetOrdersRequest) (*dto.GetOrdersResponse, error) {
	res, err := s.orderRepo.GetOrders(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}

	if len(res) == 0 {
		return nil, custom_errs.ErrNoContent
	}

	var orders []*dto.Order
	for _, order := range res {
		orders = append(orders, &dto.Order{
			Number:     order.ID,
			Status:     order.Status.ToString(),
			Accrual:    order.Accrual,
			UploadedAt: *order.UploadedAt,
		})
	}
	return &dto.GetOrdersResponse{Orders: orders}, nil
}

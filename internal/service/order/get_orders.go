package order

import (
	"context"
	"fmt"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/order/dto"
)

func (s *OrderServiceImpl) GetOrders(ctx context.Context, req *dto.GetOrdersRequest) ([]*dto.GetOrdersResponse, error) {
	s.lg.Infow("GET ORDERS REQUEST", "get_orders_request", req)

	res, err := s.orderRepo.GetOrders(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}

	if len(res) == 0 {
		return nil, custom_errs.ErrNoContent
	}

	var orders []*dto.GetOrdersResponse
	for _, order := range res {
		orders = append(orders, &dto.GetOrdersResponse{
			Number:     order.ID,
			Status:     order.Status.ToString(),
			Accrual:    order.Accrual,
			UploadedAt: *order.UploadedAt,
		})
	}
	return orders, nil
}

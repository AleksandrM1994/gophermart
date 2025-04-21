package order

import (
	"context"
	"fmt"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/order/dto"
)

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	errValidate := req.Validate()
	if errValidate != nil {
		return nil, fmt.Errorf("validate: %w", errValidate)
	}

	if !service.LunaCheck(req.Order) {
		return nil, fmt.Errorf("LunaCheck for order: %w", custom_errs.ErrWrongFormat)
	}

	order, errGetOrderByID := s.orderRepo.GetOrderByID(ctx, req.Order)
	if errGetOrderByID != nil {
		return nil, fmt.Errorf("orderRepo.GetOrderByID:%w", errGetOrderByID)
	}

	if order != nil {
		if order.ID == req.Order && order.UserID == req.UserID {
			return &dto.CreateOrderResponse{
				Order: &dto.Order{
					Number:     order.ID,
					Status:     order.Status.ToString(),
					Accrual:    order.Accrual,
					UploadedAt: *order.UploadedAt,
				},
			}, nil
		} else {
			return nil, fmt.Errorf("order already exists for other user: %w", custom_errs.ErrDuplicateKey)
		}
	}

	err := s.orderRepo.CreateOrder(ctx, &repository.Order{
		ID:         req.Order,
		Status:     repository.OrderStatusUnknown,
		UploadedAt: service.DatePtr(time.Now()),
		UserID:     req.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}
	return nil, nil
}

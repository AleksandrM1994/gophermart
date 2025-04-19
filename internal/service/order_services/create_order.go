package order_services

import (
	"context"
	"fmt"
	"time"

	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/order_services/dto"
)

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error {
	err := s.orderRepo.CreateOrder(ctx, &repository.Order{
		ID:         req.Order,
		Status:     repository.OrderStatusUnknown,
		UploadedAt: service.DatePtr(time.Now()),
	})
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}
	return nil
}

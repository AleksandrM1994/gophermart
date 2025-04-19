package order_services

import (
	"context"

	"github.com/gophermart/internal/service/order_services/dto"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error
}

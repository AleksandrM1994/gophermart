package order

import (
	"context"

	"github.com/gophermart/internal/service/order/dto"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error)
	GetOrders(ctx context.Context, req *dto.GetOrdersRequest) (*dto.GetOrdersResponse, error)
}

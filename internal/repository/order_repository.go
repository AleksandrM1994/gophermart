package repository

import (
	"context"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrders(ctx context.Context, userID string) ([]*Order, error)
	GetOrderById(ctx context.Context, id string) (*Order, error)
}

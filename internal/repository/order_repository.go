package repository

import (
	"context"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrders(ctx context.Context, userID string) ([]*Order, error)
	GetOrderByID(ctx context.Context, id string) (*Order, error)
}

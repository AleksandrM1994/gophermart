package repository

import (
	"context"
	"fmt"
)

type OrderRepositoryImpl struct {
	*Repository
}

func NewOrderRepository(repo *Repository) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{repo}
}

func (r *OrderRepositoryImpl) CreateOrder(ctx context.Context, order *Order) error {
	err := r.db.Create(order).Error
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}
	return nil
}

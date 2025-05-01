package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	*Repository
}

func NewOrderRepository(repo *Repository) OrderRepository {
	return &OrderRepositoryImpl{repo}
}

func (r *OrderRepositoryImpl) CreateOrder(ctx context.Context, order *Order) error {
	err := r.db.WithContext(ctx).Create(order).Error
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}
	return nil
}

func (r *OrderRepositoryImpl) GetOrders(ctx context.Context, userID string) ([]*Order, error) {
	var orders []*Order
	err := r.db.WithContext(ctx).Model(&Order{}).Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) GetOrderByID(ctx context.Context, id string) (*Order, error) {
	var order Order
	err := r.db.WithContext(ctx).Model(&Order{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("get orders: %w", err)
	}
	return &order, nil
}

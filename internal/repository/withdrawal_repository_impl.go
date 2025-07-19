package repository

import (
	"context"
	"fmt"
)

type WithdrawalRepositoryImpl struct {
	*Repository
}

func NewWithdrawalRepository(repository *Repository) WithdrawalRepository {
	return &WithdrawalRepositoryImpl{repository}
}

func (r *WithdrawalRepositoryImpl) CreateWithdrawal(ctx context.Context, withdrawal *Withdrawal) error {
	err := r.db.WithContext(ctx).Create(&withdrawal).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}

func (r *WithdrawalRepositoryImpl) GetWithdrawalByUserID(ctx context.Context, userID string) ([]*Withdrawal, error) {
	var withdrawals []*Withdrawal
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&withdrawals).Error
	if err != nil {
		return nil, fmt.Errorf("db.Find: %w", err)
	}
	return withdrawals, nil
}

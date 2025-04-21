package repository

import "context"

type WithdrawalRepository interface {
	CreateWithdrawal(ctx context.Context, withdrawal *Withdrawal) error
	GetWithdrawalByUserID(ctx context.Context, userID string) ([]*Withdrawal, error)
}

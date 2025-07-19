package withdrawal

import (
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
)

type WithdrawalServiceImpl struct {
	lg             *zap.SugaredLogger
	cfg            config.Config
	userRepo       repository.UserRepository
	withdrawalRepo repository.WithdrawalRepository
}

func NewService(
	lg *zap.SugaredLogger,
	cfg config.Config,
	userRepo repository.UserRepository,
	withdrawalRepo repository.WithdrawalRepository,
) *WithdrawalServiceImpl {
	return &WithdrawalServiceImpl{
		lg:             lg,
		cfg:            cfg,
		userRepo:       userRepo,
		withdrawalRepo: withdrawalRepo,
	}
}

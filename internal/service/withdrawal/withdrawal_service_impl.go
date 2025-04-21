package withdrawal

import (
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
)

type WithdrawalServiceImpl struct {
	lg             *zap.SugaredLogger
	cfg            config.Config
	userRepo       *repository.UserRepositoryImpl
	withdrawalRepo *repository.WithdrawalRepositoryImpl
}

func NewWithdrawalService(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *WithdrawalServiceImpl {
	service := &WithdrawalServiceImpl{
		lg:  lg,
		cfg: cfg,
	}
	service.userRepo = repository.NewUserRepository(repo)
	service.withdrawalRepo = repository.NewWithdrawalRepository(repo)
	return service
}

package order

import (
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/accrual"
)

type OrderServiceImpl struct {
	lg             *zap.SugaredLogger
	cfg            config.Config
	userRepository repository.UserRepository
	orderRepo      repository.OrderRepository
	accrualService accrual.AccrualService
}

func NewService(
	lg *zap.SugaredLogger,
	cfg config.Config,
	userRepository repository.UserRepository,
	orderRepo repository.OrderRepository,
	accrualService accrual.AccrualService,
) *OrderServiceImpl {
	return &OrderServiceImpl{
		lg:             lg,
		cfg:            cfg,
		userRepository: userRepository,
		orderRepo:      orderRepo,
		accrualService: accrualService,
	}
}

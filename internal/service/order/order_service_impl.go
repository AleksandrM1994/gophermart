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
	orderRepo      *repository.OrderRepositoryImpl
	accrualService *accrual.AccrualServiceImpl
}

func NewOrderService(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *OrderServiceImpl {
	service := &OrderServiceImpl{
		lg:  lg,
		cfg: cfg,
	}
	service.orderRepo = repository.NewOrderRepository(repo)
	service.accrualService = accrual.NewAccrualService(lg, cfg)
	return service
}

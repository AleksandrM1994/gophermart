package order_services

import (
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
)

type OrderServiceImpl struct {
	lg        *zap.SugaredLogger
	cfg       config.Config
	orderRepo *repository.OrderRepositoryImpl
}

func NewOrderService(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *OrderServiceImpl {
	service := &OrderServiceImpl{
		lg:  lg,
		cfg: cfg,
	}
	service.orderRepo = repository.NewOrderRepository(repo)
	return service
}

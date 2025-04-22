package accrual

import (
	"go.uber.org/zap"

	"github.com/gophermart/config"
)

type AccrualServiceImpl struct {
	lg  *zap.SugaredLogger
	cfg config.Config
}

func NewAccrualService(lg *zap.SugaredLogger, cfg config.Config) *AccrualServiceImpl {
	return &AccrualServiceImpl{
		lg:  lg,
		cfg: cfg,
	}
}

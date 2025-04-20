package balance

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
)

type BalanceController struct {
	lg   *zap.SugaredLogger
	cfg  config.Config
	repo *repository.Repository
}

func NewBalanceController(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *BalanceController {
	controller := &BalanceController{
		lg:  lg,
		cfg: cfg,
	}

	return controller
}

func (c *BalanceController) Register(r *gin.Engine) {
	//r.GET("/api/user/balance", c.GetOrders)
	//r.POST("/api/user/balance/withdraw", c.GetOrders)
	//r.GET("/api/user/withdrawals", c.CreateOrder)
}

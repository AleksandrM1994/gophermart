package withdrawal

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/withdrawal"
)

type WithdrawalController struct {
	lg                *zap.SugaredLogger
	cfg               config.Config
	withdrawalService *withdrawal.WithdrawalServiceImpl
}

func NewWithdrawalController(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *WithdrawalController {
	controller := &WithdrawalController{
		lg:  lg,
		cfg: cfg,
	}
	controller.withdrawalService = withdrawal.NewWithdrawalService(lg, cfg, repo)
	return controller
}

func (c *WithdrawalController) Register(r *gin.Engine) {
	r.GET("/api/user/balance", c.GetBalance)
	r.POST("/api/user/balance/withdraw", c.MakeWithdrawal)
	r.GET("/api/user/withdrawals", c.GetWithdrawalsInfo)
}

package withdrawal

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/handlers/middlewares"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/user"
	"github.com/gophermart/internal/service/withdrawal"
)

type WithdrawalController struct {
	lg                *zap.SugaredLogger
	cfg               config.Config
	userService       *user.UserServiceImpl
	withdrawalService *withdrawal.WithdrawalServiceImpl
}

func NewWithdrawalController(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *WithdrawalController {
	controller := &WithdrawalController{
		lg:  lg,
		cfg: cfg,
	}
	controller.userService = user.NewUserService(lg, cfg, repo)
	controller.withdrawalService = withdrawal.NewWithdrawalService(lg, cfg, repo)
	return controller
}

func (c *WithdrawalController) Register(r *gin.Engine) {
	withdrawGroup := r.Group("/api/user").Use(middlewares.Authorizer(c.lg, c.cfg, c.userService))
	withdrawGroup.GET("/balance", c.GetBalance)
	withdrawGroup.POST("/balance/withdraw", c.MakeWithdrawal)
	withdrawGroup.GET("/withdrawals", c.GetWithdrawalsInfo)
}

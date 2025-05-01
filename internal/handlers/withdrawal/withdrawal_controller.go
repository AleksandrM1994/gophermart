package withdrawal

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/middlewares"
	"github.com/gophermart/internal/service/user"
	"github.com/gophermart/internal/service/withdrawal"
)

type WithdrawalController struct {
	lg                *zap.SugaredLogger
	cfg               config.Config
	userService       *user.UserServiceImpl
	withdrawalService *withdrawal.WithdrawalServiceImpl
}

func NewWithdrawalController(
	lg *zap.SugaredLogger,
	cfg config.Config,
	userService *user.UserServiceImpl,
	withdrawalService *withdrawal.WithdrawalServiceImpl,
) *WithdrawalController {
	return &WithdrawalController{
		lg:                lg,
		cfg:               cfg,
		userService:       userService,
		withdrawalService: withdrawalService,
	}
}

func (c *WithdrawalController) RegisterRoutes(r *gin.Engine) {
	withdrawGroup := r.Group("/api/user").Use(middlewares.Authorizer(c.lg, c.cfg, c.userService))
	withdrawGroup.GET("/balance", c.GetBalance)
	withdrawGroup.POST("/balance/withdraw", c.MakeWithdrawal)
	withdrawGroup.GET("/withdrawals", c.GetWithdrawalsInfo)
}

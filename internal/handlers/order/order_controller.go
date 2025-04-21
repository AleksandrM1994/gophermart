package order

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/handlers/middlewares"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/order"
	"github.com/gophermart/internal/service/user"
)

type OrderController struct {
	lg           *zap.SugaredLogger
	cfg          config.Config
	userService  *user.UserServiceImpl
	orderService *order.OrderServiceImpl
}

func NewOrderController(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *OrderController {
	controller := &OrderController{
		lg:  lg,
		cfg: cfg,
	}
	controller.userService = user.NewUserService(lg, cfg, repo)
	controller.orderService = order.NewOrderService(lg, cfg, repo)
	return controller
}

func (c *OrderController) Register(r *gin.Engine) {
	orderGroup := r.Group("/api/user").Use(middlewares.Authorizer(c.lg, c.cfg, c.userService))
	orderGroup.POST("/orders", c.CreateOrder)
	orderGroup.GET("/orders", c.GetOrders)
}

package order

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/middlewares"
	"github.com/gophermart/internal/service/order"
	"github.com/gophermart/internal/service/user"
)

type OrderController struct {
	lg           *zap.SugaredLogger
	cfg          config.Config
	userService  *user.UserServiceImpl
	orderService *order.OrderServiceImpl
}

func NewOrderController(
	lg *zap.SugaredLogger,
	cfg config.Config,
	userService *user.UserServiceImpl,
	orderService *order.OrderServiceImpl,
) *OrderController {
	return &OrderController{
		lg:           lg,
		cfg:          cfg,
		userService:  userService,
		orderService: orderService,
	}
}

func (c *OrderController) RegisterRoutes(r *gin.Engine) {
	orderGroup := r.Group("/api/user").Use(middlewares.Authorizer(c.lg, c.cfg, c.userService))
	orderGroup.POST("/orders", c.CreateOrder)
	orderGroup.GET("/orders", c.GetOrders)
}

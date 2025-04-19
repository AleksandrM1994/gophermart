package order_handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/handlers/middlewares"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/auth_services"
	"github.com/gophermart/internal/service/order_services"
)

type OrderController struct {
	lg           *zap.SugaredLogger
	cfg          config.Config
	repo         *repository.Repository
	userService  *auth_services.UserServiceImpl
	orderService *order_services.OrderServiceImpl
}

func NewOrderController(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *OrderController {
	controller := &OrderController{
		lg:  lg,
		cfg: cfg,
	}
	controller.userService = auth_services.NewUserService(lg, cfg, repo)
	controller.orderService = order_services.NewOrderService(lg, cfg, repo)
	return controller
}

func (c *OrderController) Register(r *gin.Engine) {
	orders := r.Group("/api/user", middlewares.Authorizer(c.lg, c.cfg, c.userService))
	orders.POST("/orders", c.CreateOrder)
}

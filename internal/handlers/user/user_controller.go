package user

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/service/user"
)

type UserController struct {
	cfg     config.Config
	lg      *zap.SugaredLogger
	service *user.UserServiceImpl
}

func NewUserController(cfg config.Config, logger *zap.SugaredLogger, service *user.UserServiceImpl) *UserController {
	return &UserController{
		cfg:     cfg,
		lg:      logger,
		service: service,
	}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/user/register", c.RegisterUserHandler)
	r.POST("/api/user/login", c.AuthUserHandler)
}

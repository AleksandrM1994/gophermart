package auth_handlers

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/auth_services"
)

type UserController struct {
	cfg     config.Config
	lg      *zap.SugaredLogger
	repo    *repository.Repository
	service *auth_services.UserServiceImpl
}

func NewUserController(cfg config.Config, logger *zap.SugaredLogger, repo *repository.Repository) *UserController {
	controller := &UserController{
		cfg: cfg,
		lg:  logger,
	}
	controller.service = auth_services.NewUserService(logger, cfg, repo)
	return controller
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/user/register", c.RegisterUserHandler)
	r.POST("/api/user/login", c.AuthUserHandler)
}

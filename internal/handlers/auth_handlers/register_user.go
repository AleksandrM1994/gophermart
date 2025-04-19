package auth_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/internal/handlers/auth_handlers/api"
	"github.com/gophermart/internal/service/auth_services/dto"
)

func (c *UserController) RegisterUserHandler(ctx *gin.Context) {
	var req *api.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := c.service.CreateUser(ctx, &dto.CreateUserRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

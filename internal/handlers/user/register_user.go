package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/handlers/user/api"
	"github.com/gophermart/internal/service/user/dto"
)

func (c *UserController) RegisterUserHandler(ctx *gin.Context) {
	var req *api.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, custom_errs.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	err := c.service.CreateUser(ctx, &dto.CreateUserRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		custom_errs.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

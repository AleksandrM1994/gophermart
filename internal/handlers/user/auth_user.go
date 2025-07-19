package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/user/dto"
)

type AuthUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (c *UserController) AuthUserHandler(ctx *gin.Context) {
	var req *AuthUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, custom_errs.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	res, err := c.service.AuthUser(ctx, &dto.AuthUserRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		custom_errs.RespondWithError(ctx, err)
		return
	}
	ctx.SetCookie(
		c.cfg.AuthUserCookieName,
		res.Cookie,
		int(time.Until(*res.CookieFinish).Seconds()),
		"/",
		"",
		true,
		true,
	)
	ctx.JSON(http.StatusOK, nil)
}

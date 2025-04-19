package auth_handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/internal/handlers/auth_handlers/api"
	"github.com/gophermart/internal/service/auth_services/dto"
)

func (c *UserController) AuthUserHandler(ctx *gin.Context) {
	var req *api.AuthUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	res, err := c.service.AuthUser(ctx, &dto.AuthUserRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

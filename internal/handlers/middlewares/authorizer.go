package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/user"
	"github.com/gophermart/internal/service/user/dto"
)

func Authorizer(lg *zap.SugaredLogger, cfg config.Config, srv *user.UserServiceImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, errCookie := ctx.Request.Cookie(cfg.AuthUserCookieName)
		if errCookie != nil {
			ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: errCookie.Error(),
			})
			return
		}

		res, err := srv.CheckAuthUser(ctx, &dto.CheckAuthRequest{Cookie: cookie.Value})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			return
		}

		ctx.Set("user_id", res.UserID)

		ctx.Next()
	}
}

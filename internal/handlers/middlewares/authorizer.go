package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/user"
	"github.com/gophermart/internal/service/user/dto"
)

func Authorizer(lg *zap.SugaredLogger, cfg config.Config, srv *user.UserServiceImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		testLogin := "test"
		testPass := "test"
		authCookie, errCookie := ctx.Request.Cookie(cfg.AuthUserCookieName)
		if errCookie == nil {
			res, err := srv.CheckAuthUser(ctx, &dto.CheckAuthRequest{Cookie: authCookie.Value})
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
					Code:  http.StatusUnauthorized,
					Error: err.Error(),
				})
				return
			}
			ctx.Set("user_id", res.UserID)
		}

		resAuthUser, errAuthUser := srv.AuthUser(ctx, &dto.AuthUserRequest{
			Login:    testLogin,
			Password: testPass,
		})
		if errAuthUser != nil || resAuthUser == nil {
			resCreateUser, errCreateUser := srv.CreateUser(ctx, &dto.CreateUserRequest{
				Login:    testLogin,
				Password: testPass,
			})
			if errCreateUser != nil {
				ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
					Code:  http.StatusUnauthorized,
					Error: errCreateUser.Error(),
				})
			}

			ctx.SetCookie(
				cfg.AuthUserCookieName,
				resCreateUser.Cookie,
				int(time.Until(*resCreateUser.CookieFinish).Seconds()),
				"/",
				"",
				true,
				true,
			)
			ctx.Set("user_id", resCreateUser.UserID)
			ctx.Next()
			return
		}

		lg.Infow("auth user", "response", resAuthUser)

		ctx.SetCookie(
			cfg.AuthUserCookieName,
			resAuthUser.Cookie,
			int(time.Until(*resAuthUser.CookieFinish).Seconds()),
			"/",
			"",
			true,
			true,
		)
		ctx.Set("user_id", resAuthUser.UserID)
		ctx.Next()
	}
}

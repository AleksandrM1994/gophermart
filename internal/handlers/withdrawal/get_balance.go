package withdrawal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/withdrawal/dto"
)

type GetBalanceResponse struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

func (c *WithdrawalController) GetBalance(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, custom_errs.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: "empty user id",
		})
		return
	}

	userID := value.(string)

	res, err := c.withdrawalService.GetBalance(ctx, &dto.GetBalanceRequest{UserID: userID})
	if err != nil {
		custom_errs.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &GetBalanceResponse{
		Current:   res.Current,
		Withdrawn: res.Withdrawn,
	})
}

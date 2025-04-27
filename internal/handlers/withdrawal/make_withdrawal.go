package withdrawal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/handlers/withdrawal/api"
	"github.com/gophermart/internal/service/withdrawal/dto"
)

func (c *WithdrawalController) MakeWithdrawal(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, custom_errs.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: "empty user id",
		})
		return
	}

	userID := value.(string)

	var req *api.MakeWithdrawalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, custom_errs.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	err := c.withdrawalService.MakeWithdrawal(ctx, &dto.MakeWithdrawalRequest{
		Order:  req.Order,
		Sum:    req.Sum,
		UserID: userID,
	})
	if err != nil {
		c.lg.Errorw("Make Withdrawal Error", "make_withdrawal_error", err)
		custom_errs.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

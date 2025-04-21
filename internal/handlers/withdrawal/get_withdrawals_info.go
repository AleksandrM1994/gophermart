package withdrawal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/handlers/withdrawal/api"
	"github.com/gophermart/internal/service/withdrawal/dto"
)

func (c *WithdrawalController) GetWithdrawalsInfo(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, custom_errs.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: "empty user id",
		})
		return
	}

	userID := value.(string)

	withdrawals, err := c.withdrawalService.GetWithdrawalsInfo(ctx, &dto.GetWithdrawalsInfoRequest{
		UserID: userID,
	})
	if err != nil {
		custom_errs.RespondWithError(ctx, err)
		return
	}

	var res *api.GetWithdrawalsInfoResponse
	errCopy := copier.Copy(&res, withdrawals)
	if errCopy != nil {
		ctx.JSON(http.StatusInternalServerError, custom_errs.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: errCopy.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, withdrawals)
}

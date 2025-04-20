package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/internal/errors"
	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/order/dto"
)

func (c *OrderController) GetOrders(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: "empty user id",
		})
		return
	}

	userID := value.(string)

	res, err := c.orderService.GetOrders(ctx, &dto.GetOrdersRequest{
		UserID: userID,
	})
	if err != nil {
		custom_errs.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

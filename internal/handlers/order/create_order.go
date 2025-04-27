package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/order/dto"
)

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, custom_errs.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: "empty user id",
		})
		return
	}

	userID := value.(string)

	data, errGetRawData := ctx.GetRawData()
	if errGetRawData != nil {
		ctx.JSON(http.StatusBadRequest, custom_errs.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: errGetRawData.Error(),
		})
		return
	}

	order, errCreateOrder := c.orderService.CreateOrder(ctx, &dto.CreateOrderRequest{
		Order:  string(data),
		UserID: userID,
	})
	if errCreateOrder != nil {
		c.lg.Infow("CREATE ORDER ERROR", "create_order_error", errCreateOrder)
		custom_errs.RespondWithError(ctx, errCreateOrder)
		return
	}

	if order != nil {
		ctx.JSON(http.StatusOK, order)
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}

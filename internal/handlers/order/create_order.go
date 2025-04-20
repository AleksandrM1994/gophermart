package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/internal/errors"
	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/order/dto"
)

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: "empty user id",
		})
		return
	}

	userID := value.(string)

	data, errGetRawData := ctx.GetRawData()
	if errGetRawData != nil {
		ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: errGetRawData.Error(),
		})
		return
	}

	order, errCreateOrder := c.orderService.CreateOrder(ctx, &dto.CreateOrderRequest{
		Order:  string(data),
		UserID: userID,
	})
	if errCreateOrder != nil {
		custom_errs.RespondWithError(ctx, errCreateOrder)
		return
	}

	if order != nil {
		ctx.JSON(http.StatusOK, order)
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
	return
}

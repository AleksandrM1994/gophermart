package order_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/order_services/dto"
)

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	data, errGetRawData := ctx.GetRawData()
	if errGetRawData != nil {
		ctx.JSON(http.StatusUnauthorized, errors.ErrorResponse{
			Code:  http.StatusUnauthorized,
			Error: errGetRawData.Error(),
		})
		return
	}

	errCreateOrder := c.orderService.CreateOrder(ctx, &dto.CreateOrderRequest{
		Order: string(data),
	})
	if errCreateOrder != nil {
		ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: errCreateOrder.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

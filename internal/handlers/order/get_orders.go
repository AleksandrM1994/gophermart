package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/handlers/order/api"
	"github.com/gophermart/internal/service/order/dto"
)

func (c *OrderController) GetOrders(ctx *gin.Context) {
	value, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, custom_errs.ErrorResponse{
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
		c.lg.Errorw("GET ORDERS ERROR", "get_orders_error", err)
		custom_errs.RespondWithError(ctx, err)
		return
	}

	response := make([]*api.GetOrdersResponse, 0)
	errCopy := copier.Copy(&response, &res)
	if errCopy != nil {
		ctx.JSON(http.StatusInternalServerError, custom_errs.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: errCopy.Error(),
		})
		return
	}

	c.lg.Infow("GET ORDERS RESPONSE", "get_orders_response", response)

	ctx.JSON(http.StatusOK, response)
}

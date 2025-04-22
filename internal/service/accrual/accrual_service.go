package accrual

import (
	"github.com/gin-gonic/gin"

	"github.com/gophermart/internal/service/accrual/dto"
)

type AccrualService interface {
	GetOrderInfo(ctx *gin.Context, req *dto.GetOrderInfoRequest) (*dto.GetOrderInfoResponse, error)
}

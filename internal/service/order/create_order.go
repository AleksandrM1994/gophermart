package order

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/order/dto"
)

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	errValidate := req.Validate()
	if errValidate != nil {
		return nil, fmt.Errorf("validate: %w", errValidate)
	}

	if !LunaCheck(req.Order) {
		return nil, fmt.Errorf("LunaCheck for order: %w", custom_errs.ErrWrongFormat)
	}

	order, errGetOrderById := s.orderRepo.GetOrderById(ctx, req.Order)
	if errGetOrderById != nil {
		return nil, fmt.Errorf("orderRepo.GetOrderById:%w", errGetOrderById)
	}

	if order != nil {
		if order.ID == req.Order && order.UserID == req.UserID {
			return &dto.CreateOrderResponse{
				&dto.Order{
					Number:     order.ID,
					Status:     order.Status.ToString(),
					Accrual:    order.Accrual,
					UploadedAt: *order.UploadedAt,
				},
			}, nil
		} else {
			return nil, fmt.Errorf("order already exists for other user: %w", custom_errs.ErrDuplicateKey)
		}
	}

	err := s.orderRepo.CreateOrder(ctx, &repository.Order{
		ID:         req.Order,
		Status:     repository.OrderStatusUnknown,
		UploadedAt: service.DatePtr(time.Now()),
		UserID:     req.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}
	return nil, nil
}

func LunaCheck(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	if len(number) == 0 {
		return false
	}

	sum := 0
	shouldDouble := false

	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}

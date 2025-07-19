package order

import (
	"context"
	"fmt"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	accrual_dto "github.com/gophermart/internal/service/accrual/dto"
	"github.com/gophermart/internal/service/order/dto"
)

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	s.lg.Infow("CREATE ORDER REQUEST", "create_order_request", req)

	errValidate := req.Validate()
	if errValidate != nil {
		return nil, fmt.Errorf("validate: %w", errValidate)
	}

	if !service.LunaCheck(req.Order) {
		return nil, fmt.Errorf("LunaCheck for order: %w", custom_errs.ErrWrongFormat)
	}

	order, errGetOrderByID := s.orderRepo.GetOrderByID(ctx, req.Order)
	if errGetOrderByID != nil {
		return nil, fmt.Errorf("orderRepo.GetOrderByID:%w", errGetOrderByID)
	}

	if order != nil {
		if order.ID == req.Order && order.UserID == req.UserID {
			return &dto.CreateOrderResponse{
				Order: &dto.Order{
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

	res, errGetOrderInfo := s.accrualService.GetOrderInfo(ctx, &accrual_dto.GetOrderInfoRequest{
		Order: req.Order,
	})
	if errGetOrderInfo != nil {
		return nil, fmt.Errorf("accrualService.GetOrderInfo:%w", errGetOrderInfo)
	}

	mosLoc, errLoadLocation := time.LoadLocation("Europe/Moscow")
	if errLoadLocation != nil {
		return nil, fmt.Errorf("time.LoadLocation:%w", errLoadLocation)
	}
	uploadAtString := time.Now().In(mosLoc).Format(time.RFC3339)
	uploadAt, errTimeParse := time.Parse(time.RFC3339, uploadAtString)
	if errTimeParse != nil {
		return nil, fmt.Errorf("time.Parse:%w", errTimeParse)
	}

	err := s.orderRepo.CreateOrder(ctx, &repository.Order{
		ID:         req.Order,
		Status:     repository.OrderStatus(res.Status),
		Accrual:    res.Accrual,
		UploadedAt: service.DatePtr(uploadAt),
		UserID:     req.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}

	if res.Status == repository.OrderStatusProcessed.ToString() {
		errUpdateUserByID := s.userRepository.UpdateUserByID(ctx, req.UserID, func(currentUser *repository.User) error {
			currentUser.Balance += res.Accrual
			return nil
		})
		if errUpdateUserByID != nil {
			return nil, fmt.Errorf("userRepository.UpdateUserByID:%w", errUpdateUserByID)
		}
	}

	return nil, nil
}

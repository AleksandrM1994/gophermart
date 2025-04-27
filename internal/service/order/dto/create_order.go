package dto

import (
	"fmt"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
)

type CreateOrderRequest struct {
	Order  string
	UserID string
}

func (r CreateOrderRequest) Validate() error {
	switch {
	case r.Order == "":
		return fmt.Errorf("order is required:%w", custom_errs.ErrValidate)
	case r.UserID == "":
		return fmt.Errorf("user id is required:%w", custom_errs.ErrValidate)
	}
	return nil
}

type CreateOrderResponse struct {
	Order *Order
}

type Order struct {
	Number     string
	Status     string
	Accrual    float32
	UploadedAt time.Time
}

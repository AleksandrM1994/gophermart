package accrual

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/accrual/dto"
)

func (s *AccrualServiceImpl) GetOrderInfo(ctx context.Context, req *dto.GetOrderInfoRequest) (*dto.GetOrderInfoResponse, error) {
	s.lg.Infow("ACCRUAL GET ORDER INFO REQUEST", "get_orders_info_request", req)

	uri := fmt.Sprintf("/api/orders/%s", req.Order)
	url := s.cfg.AccrualSystemAddress + uri

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	s.lg.Infow("ACCRUAL GET ORDER INFO RESPONSE", "get_orders_info_status_code", resp.StatusCode, "get_orders_info_body", resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		var res *dto.GetOrderInfoResponse
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		s.lg.Infow("get order info", "response", res)
		return res, nil
	case http.StatusNoContent:
		return &dto.GetOrderInfoResponse{
			Order:  req.Order,
			Status: repository.OrderStatusNew.ToString(),
		}, nil
	case http.StatusTooManyRequests:
		return nil, custom_errs.ErrManyRequests
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error")
	default:
		body, errReadAll := io.ReadAll(resp.Body)
		if errReadAll != nil {
			return nil, fmt.Errorf("io.ReadAll: %w", errReadAll)
		}
		return nil, fmt.Errorf("internal server error: %s", string(body))
	}
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	accrual_dto "github.com/gophermart/internal/service/accrual/dto"
)

func (suite *EndpointsTestSuite) Test_CreateOrderHandler_Success(t *testing.T) {
	suite.userRepo.EXPECT().GetUserByID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return(getTestUser(), nil).Times(1)

	suite.orderRepo.EXPECT().GetOrderByID(gomock.Any(), "12345678903").Return(
		nil, nil).Times(1)

	suite.accrualService.EXPECT().GetOrderInfo(gomock.Any(), &accrual_dto.GetOrderInfoRequest{
		Order: "12345678903",
	}).Return(&accrual_dto.GetOrderInfoResponse{
		Order:   "12345678903",
		Status:  repository.OrderStatusNew.ToString(),
		Accrual: 0,
	}, nil).Times(1)

	mosLoc, _ := time.LoadLocation("Europe/Moscow")
	uploadAtString := time.Now().In(mosLoc).Format(time.RFC3339)
	uploadAt, _ := time.Parse(time.RFC3339, uploadAtString)

	suite.orderRepo.EXPECT().CreateOrder(gomock.Any(), &repository.Order{
		ID:         "12345678903",
		Accrual:    0,
		Status:     repository.OrderStatusNew,
		UploadedAt: service.DatePtr(uploadAt),
		UserID:     getTestUser().ID,
	}).Return(nil).Times(1)

	reqBody := `12345678903`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(reqBody))
	req.AddCookie(getTestCookie())
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Errorf("expected status code %d, got %d", http.StatusAccepted, w.Code)
	}
}

func (suite *EndpointsTestSuite) Test_CreateOrderHandler_Unauthorized(t *testing.T) {
	reqBody := `12345678903`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func (suite *EndpointsTestSuite) Test_CreateOrderHandler_ValidateError(t *testing.T) {
	suite.userRepo.EXPECT().GetUserByID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return(getTestUser(), nil).Times(1)

	reqBody := ``
	req, _ := http.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(reqBody))
	req.AddCookie(getTestCookie())
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func (suite *EndpointsTestSuite) Test_CreateOrderHandler_WrongOrderFormat(t *testing.T) {
	suite.userRepo.EXPECT().GetUserByID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return(getTestUser(), nil).Times(1)

	reqBody := `12345`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(reqBody))
	req.AddCookie(getTestCookie())
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}
}

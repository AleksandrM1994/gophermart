package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
)

func (suite *EndpointsTestSuite) Test_GetOrdersHandler(t *testing.T) {
	suite.userRepo.EXPECT().GetUserByID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return(getTestUser(), nil).Times(1)

	suite.orderRepo.EXPECT().GetOrders(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return([]*repository.Order{
			{
				ID:         "12345678903",
				Status:     repository.OrderStatusNew,
				Accrual:    0,
				UploadedAt: service.DatePtr(time.Now()),
				UserID:     getTestUser().ID,
			},
		}, nil).Times(1)

	req, _ := http.NewRequest(http.MethodGet, "/api/user/orders", nil)
	req.AddCookie(getTestCookie())
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

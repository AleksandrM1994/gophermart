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

func (suite *EndpointsTestSuite) Test_GetWithdrawalsHandler(t *testing.T) {
	suite.userRepo.EXPECT().GetUserByID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return(getTestUser(), nil).Times(1)

	suite.withdrawalRepo.EXPECT().GetWithdrawalByUserID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return([]*repository.Withdrawal{
			{
				OrderID:     "2377225624",
				Sum:         751,
				ProcessedAt: service.DatePtr(time.Now()),
				UserID:      getTestUser().ID,
			},
		}, nil).Times(1)

	req, _ := http.NewRequest(http.MethodGet, "/api/user/withdrawals", nil)
	req.AddCookie(getTestCookie())
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

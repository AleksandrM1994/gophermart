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
)

func (suite *EndpointsTestSuite) Test_MakeWithdrawalHandler(t *testing.T) {
	userWithBalance := getTestUser()
	userWithBalance.Balance = 1000

	suite.userRepo.EXPECT().GetUserByID(gomock.Any(), "6d014812-cb3e-44a1-b3be-701b1f5bdb87").
		Return(userWithBalance, nil).Times(2)

	mosLoc, _ := time.LoadLocation("Europe/Moscow")
	processAtString := time.Now().In(mosLoc).Format(time.RFC3339)
	processAt, _ := time.Parse(time.RFC3339, processAtString)

	suite.withdrawalRepo.EXPECT().CreateWithdrawal(
		gomock.Any(),
		&repository.Withdrawal{
			OrderID:     "2377225624",
			Sum:         751,
			ProcessedAt: service.DatePtr(processAt),
			UserID:      getTestUser().ID,
		},
	).Return(nil).Times(1)

	suite.userRepo.EXPECT().UpdateUserByID(
		gomock.Any(),
		"6d014812-cb3e-44a1-b3be-701b1f5bdb87",
		gomock.Any(),
	).Return(nil).Times(1)

	reqBody := `{"order": "2377225624", "sum": 751}`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/balance/withdraw", strings.NewReader(reqBody))
	req.AddCookie(getTestCookie())
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusAccepted, w.Code)
	}
}

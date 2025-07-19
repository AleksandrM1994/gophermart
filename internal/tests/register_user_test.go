package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func (suite *EndpointsTestSuite) Test_RegisterUserHandler_Success(t *testing.T) {
	//hashData, _ := service.HashData(suite.cfg.HashSecret, []byte("test"))

	suite.userRepo.EXPECT().CreateUser(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(nil).Times(1)

	reqBody := `{"login": "test", "password": "test"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func (suite *EndpointsTestSuite) Test_RegisterUserHandler_ValidateError(t *testing.T) {
	reqBody := `{"login": "test", "password": ""}`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

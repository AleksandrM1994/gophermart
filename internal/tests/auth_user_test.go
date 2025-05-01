package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gophermart/internal/service"
)

func (suite *EndpointsTestSuite) Test_AuthUserHandler_Success(t *testing.T) {
	hashData, _ := service.HashData(suite.cfg.HashSecret, []byte("test"))

	suite.userRepo.EXPECT().GetUserByLogPass(
		gomock.Any(),
		hashData,
		hashData,
	).Return(getTestUser(), nil).Times(1)

	reqBody := `{"login": "test", "password": "test"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func (suite *EndpointsTestSuite) Test_AuthUserHandler_Unauthorized(t *testing.T) {
	hashData, _ := service.HashData(suite.cfg.HashSecret, []byte("test"))

	suite.userRepo.EXPECT().GetUserByLogPass(
		gomock.Any(),
		hashData,
		hashData,
	).Return(nil, nil).Times(1)

	reqBody := `{"login": "test", "password": "test"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func (suite *EndpointsTestSuite) Test_AuthUserHandler_ValidateError(t *testing.T) {
	reqBody := `{"login": "test", "password": ""}`
	req, _ := http.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

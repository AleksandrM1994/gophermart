package auth_services

import (
	"context"
	"fmt"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service/auth_services/dto"
)

func (s *UserServiceImpl) CheckAuthUser(ctx context.Context, req *dto.CheckAuthRequest) (*dto.CheckAuthResponse, error) {
	var userID string
	err := s.cookie.Decode(s.cfg.AuthUserCookieName, req.Cookie, &userID)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cookie: %w", err)
	}

	if userID == "" {
		return nil, custom_errs.ErrUnauthorized
	}

	user, errGetUserByID := s.userRepo.GetUserByID(ctx, userID)
	if errGetUserByID != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", errGetUserByID)
	}

	session := user.Session

	if time.Now().After(*session.CookieFinish) || session.Cookie != req.Cookie {
		return nil, custom_errs.ErrUnauthorized
	}

	return &dto.CheckAuthResponse{UserID: userID}, nil
}

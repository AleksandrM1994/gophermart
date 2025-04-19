package auth_services

import (
	"context"
	"fmt"

	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/auth_services/dto"
)

func (s *UserServiceImpl) AuthUser(ctx context.Context, req *dto.AuthUserRequest) (*dto.AuthUserResponse, error) {
	loginHash, err := service.HashData(s.cfg.HashSecret, []byte(req.Login))
	if err != nil {
		return nil, fmt.Errorf("failed to hash login: %w", err)
	}

	passHash, err := service.HashData(s.cfg.HashSecret, []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := s.userRepo.GetUserByLogPass(ctx, loginHash, passHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &dto.AuthUserResponse{Cookie: user.Session.Cookie, CookieFinish: user.Session.CookieFinish}, nil
}

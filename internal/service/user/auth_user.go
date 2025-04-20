package user

import (
	"context"
	"fmt"

	custom_errs "github.com/gophermart/internal/errors"
	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/user/dto"
)

func (s *UserServiceImpl) AuthUser(ctx context.Context, req *dto.AuthUserRequest) (*dto.AuthUserResponse, error) {
	errValidate := req.Validate()
	if errValidate != nil {
		return nil, fmt.Errorf("validate: %w", errValidate)
	}

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

	if user == nil {
		return nil, fmt.Errorf("user not found: %w", custom_errs.ErrUnauthorized)
	}

	return &dto.AuthUserResponse{Cookie: user.Session.Cookie, CookieFinish: user.Session.CookieFinish}, nil
}

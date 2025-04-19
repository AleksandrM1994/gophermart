package auth_services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
	"github.com/gophermart/internal/service/auth_services/dto"
)

func (s *UserServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserRequest) error {
	loginHash, err := service.HashData(s.cfg.HashSecret, []byte(req.Login))
	if err != nil {
		return fmt.Errorf("failed to hash login: %w", err)
	}

	passHash, err := service.HashData(s.cfg.HashSecret, []byte(req.Password))
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	id := uuid.New().String()

	cookie, err := s.cookie.Encode(s.cfg.AuthUserCookieName, id)
	if err != nil {
		return fmt.Errorf("failed to encode cookie: %w", err)
	}

	cookieFinish := service.DatePtr(time.Now().Add(24 * time.Hour))

	errCreateUser := s.userRepo.CreateUser(
		ctx,
		&repository.User{
			ID:       id,
			Login:    loginHash,
			Password: passHash,
		},
		&repository.Session{
			Cookie:       cookie,
			CookieFinish: cookieFinish,
			UserID:       id,
		})
	if errCreateUser != nil {
		return fmt.Errorf("failed to create user: %w", errCreateUser)
	}

	return nil
}

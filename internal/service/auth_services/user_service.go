package auth_services

import (
	"context"

	dto2 "github.com/gophermart/internal/service/auth_services/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto2.CreateUserRequest) error
	AuthUser(ctx context.Context, req *dto2.CreateUserRequest) (*dto2.AuthUserResponse, error)
	CheckAuthUser(ctx context.Context, req *dto2.CheckAuthRequest) (*dto2.CheckAuthResponse, error)
}

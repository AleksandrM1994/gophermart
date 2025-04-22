package user

import (
	"context"

	dto "github.com/gophermart/internal/service/user/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	AuthUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.AuthUserResponse, error)
	CheckAuthUser(ctx context.Context, req *dto.CheckAuthRequest) (*dto.CheckAuthResponse, error)
}

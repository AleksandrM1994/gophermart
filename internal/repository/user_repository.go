package repository

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User, session *Session) error
	GetUserByLogPass(ctx context.Context, login, password string) (*User, error)
	GetUserByID(ctx context.Context, UserID string) (*User, error)
	UpdateUserByID(ctx context.Context, userID string, updateFunc func(user *User) error) error
}

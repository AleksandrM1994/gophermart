package repository

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User, session *Session) error
	GetSessionByLogPass(ctx context.Context, login, password string) (*Session, error)
	GetUserByID(ctx context.Context, UserID string) (*User, error)
}

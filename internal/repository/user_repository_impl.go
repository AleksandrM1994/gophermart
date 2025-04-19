package repository

import (
	"context"
	"fmt"
)

type UserRepositoryImpl struct {
	*Repository
}

func NewUserRepository(repo *Repository) *UserRepositoryImpl {
	return &UserRepositoryImpl{repo}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *User, session *Session) error {
	err := r.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = r.db.Create(session).Error
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByLogPass(ctx context.Context, login, password string) (*User, error) {
	var user *User
	err := r.db.Model(&user).Where("login = ? and password = ?", login, password).Preload("Session").Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, userID string) (*User, error) {
	var user *User
	err := r.db.Model(&user).Where("id = ?", userID).Preload("Session").Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

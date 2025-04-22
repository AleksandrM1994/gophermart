package dto

import (
	"fmt"
	"time"

	custom_errs "github.com/gophermart/internal/errors"
)

type CreateUserRequest struct {
	Login    string
	Password string
}

func (r *CreateUserRequest) Validate() error {
	switch {
	case r.Login == "":
		return fmt.Errorf("invalid login:%w", custom_errs.ErrValidate)
	case r.Password == "":
		return fmt.Errorf("invalid password:%w", custom_errs.ErrValidate)
	}
	return nil
}

type CreateUserResponse struct {
	Cookie       string
	CookieFinish *time.Time
	UserID       string
}

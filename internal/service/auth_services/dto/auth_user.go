package dto

import "time"

type AuthUserRequest struct {
	Login    string
	Password string
}

type AuthUserResponse struct {
	Cookie       string
	CookieFinish *time.Time
}

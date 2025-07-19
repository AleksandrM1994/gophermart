package dto

type CheckAuthRequest struct {
	Cookie string
}

type CheckAuthResponse struct {
	UserID string
}

package api

type AuthUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

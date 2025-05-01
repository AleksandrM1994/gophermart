package tests

import (
	"net/http"
	"time"

	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service"
)

//hashData, _ = service.HashData("my_secret", []byte("test"))

func getTestUser() *repository.User {
	return &repository.User{
		ID:       "6d014812-cb3e-44a1-b3be-701b1f5bdb87",
		Login:    "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
		Password: "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
		Session: repository.Session{
			Cookie:       "MTc0NTc4NDMyMXxNazQwQTZBNlVVVkEzbGE4OVlGak9TQUVWS1ViaHhhbVBoLXZVSVJReEVKcW5kVDA5eXh2V19jeV9rVzBtY2xDTlhRQXAzVFNveFk9fLV7v7hqzn007xv1Vxa1NLY8Af0bOWCoCnQEovceTq5k",
			CookieFinish: service.DatePtr(time.Now().Add(24 * time.Hour)),
			UserID:       "6d014812-cb3e-44a1-b3be-701b1f5bdb87",
		},
	}
}

func getTestCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "Authorization",
		Value:   getTestUser().Session.Cookie,
		Expires: *getTestUser().Session.CookieFinish,
		Path:    "/",
	}
}

package user

import (
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/repository"
)

type UserServiceImpl struct {
	lg       *zap.SugaredLogger
	cfg      config.Config
	userRepo repository.UserRepository
	cookie   *securecookie.SecureCookie
}

func NewService(lg *zap.SugaredLogger, cfg config.Config, userRepo repository.UserRepository) *UserServiceImpl {
	srv := &UserServiceImpl{lg: lg, cfg: cfg, userRepo: userRepo}
	srv.cookie = newSecureCookie()
	return srv
}

func newSecureCookie() *securecookie.SecureCookie {
	var hashKey = []byte("very-very-very-very-secret-key32")
	var blockKey = []byte("a-lot-of-secret!")
	return securecookie.New(hashKey, blockKey)
}

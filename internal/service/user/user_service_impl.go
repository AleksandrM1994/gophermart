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
	userRepo *repository.UserRepositoryImpl
	cookie   *securecookie.SecureCookie
}

func NewUserService(lg *zap.SugaredLogger, cfg config.Config, repo *repository.Repository) *UserServiceImpl {
	srv := &UserServiceImpl{lg: lg, cfg: cfg}
	srv.userRepo = repository.NewUserRepository(repo)
	srv.cookie = newSecureCookie()
	return srv
}

func newSecureCookie() *securecookie.SecureCookie {
	var hashKey = []byte("very-very-very-very-secret-key32")
	var blockKey = []byte("a-lot-of-secret!")
	return securecookie.New(hashKey, blockKey)
}

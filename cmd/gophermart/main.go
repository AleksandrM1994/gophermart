package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	"github.com/gophermart/internal/handlers/order"
	"github.com/gophermart/internal/handlers/user"
	"github.com/gophermart/internal/repository"
)

func main() {
	logger, loggerErr := zap.NewDevelopment()
	if loggerErr != nil {
		panic("cannot initialize zap")
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	lg := logger.Sugar()

	cfg := config.Init()

	g := gin.Default()

	repo, err := repository.NewRepository(cfg, lg)
	if err != nil {
		lg.Fatalf("repository.NewRepository, %w", err)
	}

	userController := user.NewUserController(cfg, lg, repo)
	userController.RegisterRoutes(g)

	orderController := order.NewOrderController(lg, cfg, repo)
	orderController.Register(g)

	err = g.Run(cfg.HTTPAddress)
	if err != nil {
		lg.Fatalf("g.Run, %w", err)
	}
}

package main

import (
	"net/http"
	"time"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	orderHandlers "github.com/gophermart/internal/handlers/order"
	userHandlers "github.com/gophermart/internal/handlers/user"
	withdrawalHandlers "github.com/gophermart/internal/handlers/withdrawal"
	"github.com/gophermart/internal/repository"
	"github.com/gophermart/internal/service/accrual"
	orderService "github.com/gophermart/internal/service/order"
	userService "github.com/gophermart/internal/service/user"
	withdrawalService "github.com/gophermart/internal/service/withdrawal"
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

	lg := *logger.Sugar()

	cfg, errNewConfig := config.NewConfig()
	if errNewConfig != nil {
		panic(errNewConfig)
	}

	g := gin.Default()

	repo, err := repository.NewRepository(cfg, &lg)
	if err != nil {
		lg.Fatalf("repository.NewRepository, %w", err)
	}

	userRepo := repository.NewUserRepository(repo)
	orderRepo := repository.NewOrderRepository(repo)
	withdrawalRepo := repository.NewWithdrawalRepository(repo)

	accrualService := accrual.NewAccrualService(&lg, cfg)
	userServiceImpl := userService.NewService(&lg, cfg, userRepo)
	orderServiceImpl := orderService.NewService(&lg, cfg, userRepo, orderRepo, accrualService)
	withdrawalServiceImpl := withdrawalService.NewService(&lg, cfg, userRepo, withdrawalRepo)

	userController := userHandlers.NewUserController(cfg, &lg, userServiceImpl)
	userController.RegisterRoutes(g)

	orderController := orderHandlers.NewOrderController(&lg, cfg, userServiceImpl, orderServiceImpl)
	orderController.RegisterRoutes(g)

	withdrawalController := withdrawalHandlers.NewWithdrawalController(&lg, cfg, userServiceImpl, withdrawalServiceImpl)
	withdrawalController.RegisterRoutes(g)

	server := &http.Server{
		Addr:         cfg.HTTPAddress,
		Handler:      g,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	err = server.ListenAndServe()
	if err != nil {
		lg.Fatalf("g.Run, %w", err)
	}
}

package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"github.com/gophermart/config"
	orderHandlers "github.com/gophermart/internal/handlers/order"
	userHandlers "github.com/gophermart/internal/handlers/user"
	withdrawalHandlers "github.com/gophermart/internal/handlers/withdrawal"
	"github.com/gophermart/internal/mocks"
	"github.com/gophermart/internal/service/order"
	"github.com/gophermart/internal/service/user"
	"github.com/gophermart/internal/service/withdrawal"
)

type EndpointsTestSuite struct {
	cfg                  config.Config
	userController       *userHandlers.UserController
	orderController      *orderHandlers.OrderController
	withdrawalController *withdrawalHandlers.WithdrawalController
	accrualService       *mocks.MockAccrualService
	orderRepo            *mocks.MockOrderRepository
	userRepo             *mocks.MockUserRepository
	withdrawalRepo       *mocks.MockWithdrawalRepository
	router               *gin.Engine
}

func TestEndpointSuite(t *testing.T) {
	logger, loggerErr := zap.NewDevelopment()
	if loggerErr != nil {
		panic("cannot initialize zap")
	}
	lg := logger.Sugar()
	cfg, errNewConfig := config.NewConfig()
	if errNewConfig != nil {
		panic(errNewConfig)
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	orderRepo := mocks.NewMockOrderRepository(ctrl)
	withdrawalRepo := mocks.NewMockWithdrawalRepository(ctrl)
	accrualService := mocks.NewMockAccrualService(ctrl)

	userServiceImpl := user.NewService(lg, cfg, userRepo)
	orderServiceImpl := order.NewService(lg, cfg, userRepo, orderRepo, accrualService)
	withdrawalServiceImpl := withdrawal.NewService(lg, cfg, userRepo, withdrawalRepo)

	// Инициализация контроллеров
	userController := userHandlers.NewUserController(cfg, lg, userServiceImpl)
	userController.RegisterRoutes(router)

	orderController := orderHandlers.NewOrderController(lg, cfg, userServiceImpl, orderServiceImpl)
	orderController.RegisterRoutes(router)

	withdrawalController := withdrawalHandlers.NewWithdrawalController(lg, cfg, userServiceImpl, withdrawalServiceImpl)
	withdrawalController.RegisterRoutes(router)

	suite := &EndpointsTestSuite{
		cfg:                  cfg,
		userController:       userController,
		orderController:      orderController,
		withdrawalController: withdrawalController,
		accrualService:       accrualService,
		orderRepo:            orderRepo,
		userRepo:             userRepo,
		withdrawalRepo:       withdrawalRepo,
		router:               router,
	}

	suite.Test_RegisterUserHandler_Success(t)
	suite.Test_RegisterUserHandler_ValidateError(t)

	suite.Test_AuthUserHandler_Success(t)
	suite.Test_AuthUserHandler_Unauthorized(t)
	suite.Test_AuthUserHandler_ValidateError(t)

	suite.Test_CreateOrderHandler_Success(t)
	suite.Test_CreateOrderHandler_Unauthorized(t)
	suite.Test_CreateOrderHandler_ValidateError(t)
	suite.Test_CreateOrderHandler_WrongOrderFormat(t)

	suite.Test_GetOrdersHandler(t)
	suite.Test_MakeWithdrawalHandler(t)
	suite.Test_GetBalanceHandler(t)
	suite.Test_GetWithdrawalsHandler(t)
}

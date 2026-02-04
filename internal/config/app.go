package config

import (
	"xyz-multifinance/internal/delivery/http"
	"xyz-multifinance/internal/delivery/http/middleware"
	"xyz-multifinance/internal/delivery/http/route"
	"xyz-multifinance/internal/repository"
	"xyz-multifinance/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	customerRepository := repository.NewCustomerRepository(config.Log)
	transactionRepository := repository.NewTransactionRepository(config.Log)

	// setup use cases
	customerUseCase := usecase.NewCustomerUseCase(config.DB, config.Log, customerRepository)
	transactionUseCase := usecase.NewTransactionUseCase(config.DB, config.Log, transactionRepository)

	// setup controller
	customerController := http.NewCustomerController(customerUseCase, config.Log, config.Validate)
	transactionController := http.NewTransactionController(transactionUseCase, config.Log, config.Validate)

	// setup middleware
	authMiddleware := middleware.NewAuth(customerUseCase)

	routeConfig := route.RouteConfig{
		App:                   config.App,
		CustomerController:    customerController,
		TransactionController: transactionController,
		AuthMiddleware:        authMiddleware,
	}
	routeConfig.Setup()
}

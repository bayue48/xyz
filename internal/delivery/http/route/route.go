package route

import (
	"time"
	"xyz-multifinance/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type RouteConfig struct {
	App                   *fiber.App
	CustomerController    *http.CustomerController
	TransactionController *http.TransactionController
	AuthMiddleware        fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
	}))
	c.App.Post("/api/v1/users", c.CustomerController.Register)
	c.App.Post("/api/v1/users/_login", c.CustomerController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/v1/users/_profile", c.CustomerController.GetProfile)
	c.App.Post("/api/v1/users/_logout", c.CustomerController.Logout)

	c.App.Post("/api/v1/transactions", c.TransactionController.CreateTransaction)
	c.App.Get("/api/v1/transactions", c.TransactionController.GetAllTransactions)
}

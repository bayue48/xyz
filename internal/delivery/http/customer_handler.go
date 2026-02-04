package http

import (
	"net/http"
	"xyz-multifinance/internal/delivery/http/middleware"
	"xyz-multifinance/internal/model/request"
	"xyz-multifinance/internal/usecase"
	"xyz-multifinance/pkg/errors"
	http_r "xyz-multifinance/pkg/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	UseCase  *usecase.CustomerUseCase
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewCustomerController(useCase *usecase.CustomerUseCase, log *logrus.Logger, validate *validator.Validate) *CustomerController {
	return &CustomerController{
		Log:      log,
		UseCase:  useCase,
		Validate: validate,
	}
}

func (c *CustomerController) Register(ctx *fiber.Ctx) (err error) {
	req := request.CustomerReq{}
	err = http_r.BindRequest(ctx, c.Validate, &req)
	if err != nil {
		return http_r.SendError(ctx, errors.ValidationErrorToAppError(err))
	}

	err = c.UseCase.CreateCustomer(ctx.UserContext(), &req)
	if err != nil {
		return http_r.SendError(ctx, err)
	}
	return http_r.SendSuccess(ctx, http.StatusCreated, "", nil)
}

func (c *CustomerController) Login(ctx *fiber.Ctx) (err error) {
	req := request.LoginReq{}
	err = http_r.BindRequest(ctx, c.Validate, &req)
	if err != nil {
		return http_r.SendError(ctx, errors.ValidationErrorToAppError(err))
	}

	token, err := c.UseCase.LoginCustomer(ctx.UserContext(), req.NIK)
	if err != nil {
		return http_r.SendError(ctx, err)
	}
	return http_r.SendSuccess(ctx, http.StatusOK, "Login successful", token)
}

func (c *CustomerController) GetProfile(ctx *fiber.Ctx) (err error) {
	auth := middleware.GetUser(ctx)

	customer, err := c.UseCase.CurrentCustomer(ctx.UserContext(), auth.ID)
	if err != nil {
		return http_r.SendError(ctx, err)
	}
	return http_r.SendSuccess(ctx, http.StatusOK, "Profile retrieved successfully", customer)
}

func (c *CustomerController) Logout(ctx *fiber.Ctx) (err error) {
	auth := middleware.GetUser(ctx)

	err = c.UseCase.LogoutCustomer(ctx.UserContext(), auth.ID)
	if err != nil {
		return http_r.SendError(ctx, err)
	}
	return http_r.SendSuccess(ctx, http.StatusOK, "Logout successful", nil)
}

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

type TransactionController struct {
	UseCase  *usecase.TransactionUseCase
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewTransactionController(useCase *usecase.TransactionUseCase, log *logrus.Logger, validate *validator.Validate) *TransactionController {
	return &TransactionController{
		Log:      log,
		UseCase:  useCase,
		Validate: validate,
	}
}

func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) (err error) {
	auth := middleware.GetUser(ctx)
	req := request.TransactionReq{}
	err = http_r.BindRequest(ctx, c.Validate, &req)
	if err != nil {
		c.Log.Errorf("Bind Request Error : %+v", err)
		return http_r.SendError(ctx, errors.ValidationErrorToAppError(err))
	}
	req.CustomerID = auth.ID
	err = c.UseCase.CreateTransaction(ctx.UserContext(), &req)
	if err != nil {
		c.Log.Errorf("Create Transaction Error : %+v", err)
		return http_r.SendError(ctx, err)
	}
	return http_r.SendSuccess(ctx, http.StatusCreated, "Transaction created successfully", nil)
}

func (c *TransactionController) GetAllTransactions(ctx *fiber.Ctx) (err error) {
	auth := middleware.GetUser(ctx)
	transactions, err := c.UseCase.GetAllTransactions(ctx.UserContext(), auth.ID)
	if err != nil {
		return http_r.SendError(ctx, err)
	}
	return http_r.SendSuccess(ctx, http.StatusOK, "Transactions retrieved successfully", transactions)
}

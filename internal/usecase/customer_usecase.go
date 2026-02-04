package usecase

import (
	"context"
	"xyz-multifinance/internal/entity"
	"xyz-multifinance/internal/model/request"
	"xyz-multifinance/internal/model/response"
	"xyz-multifinance/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	CustomerRepository *repository.CustomerRepository
}

func NewCustomerUseCase(db *gorm.DB, logger *logrus.Logger,
	CustomerRepository *repository.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                 db,
		Log:                logger,
		CustomerRepository: CustomerRepository,
	}
}

func (uc *CustomerUseCase) CreateCustomer(ctx context.Context, customer *request.CustomerReq) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := uc.CustomerRepository.FindByNik(tx, &entity.Customer{}, customer.NIK)
	if err == nil {
		uc.Log.Error("customer with given NIK already exists")
		return fiber.ErrConflict
	}

	err = uc.CustomerRepository.Create(tx, customer)
	if err != nil {
		uc.Log.Error(err)
		return fiber.ErrInternalServerError
	}
	tx.Commit()
	return nil
}

func (uc *CustomerUseCase) LoginCustomer(ctx context.Context, nik string) (*response.LoginRes, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	var customer entity.Customer
	err := uc.CustomerRepository.FindByNik(tx, &customer, nik)
	if err != nil {
		uc.Log.Error(err)
		return nil, fiber.ErrUnauthorized
	}

	customer.Token = uuid.New().String()
	err = uc.CustomerRepository.Update(tx, &customer)
	if err != nil {
		uc.Log.Error(err)
		return nil, fiber.ErrInternalServerError
	}
	tx.Commit()
	return &response.LoginRes{Token: customer.Token}, nil
}

func (uc *CustomerUseCase) CurrentCustomer(ctx context.Context, token int64) (*response.CustomerRes, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	var customer entity.Customer
	err := uc.CustomerRepository.FindByID(tx, &customer, token)
	if err != nil {
		uc.Log.Error(err)
		return nil, fiber.ErrUnauthorized
	}
	tx.Commit()
	return response.ToCustomerResponse(&customer), nil
}

func (uc *CustomerUseCase) LogoutCustomer(ctx context.Context, token int64) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	var customer entity.Customer
	err := uc.CustomerRepository.FindByID(tx, &customer, token)
	if err != nil {
		uc.Log.Error(err)
		return fiber.ErrUnauthorized
	}

	customer.Token = ""
	err = uc.CustomerRepository.Update(tx, &customer)
	if err != nil {
		uc.Log.Error(err)
		return fiber.ErrInternalServerError
	}
	tx.Commit()
	return nil
}

func (uc *CustomerUseCase) VerifyToken(ctx context.Context, token string) (*response.Auth, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	var customer entity.Customer
	err := uc.CustomerRepository.FindByToken(tx, &customer, token)
	if err != nil {
		uc.Log.Error(err)
		return nil, fiber.ErrUnauthorized
	}
	tx.Commit()
	return &response.Auth{ID: customer.ID}, nil
}

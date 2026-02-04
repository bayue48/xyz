package usecase

import (
	"context"
	"fmt"
	"strings"
	"xyz-multifinance/internal/entity"
	"xyz-multifinance/internal/model/request"
	"xyz-multifinance/internal/model/response"
	"xyz-multifinance/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	TransactionRepository *repository.TransactionRepository
}

func NewTransactionUseCase(db *gorm.DB, logger *logrus.Logger,
	TransactionRepository *repository.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{
		DB:                    db,
		Log:                   logger,
		TransactionRepository: TransactionRepository,
	}
}

func (uc *TransactionUseCase) CreateTransaction(ctx context.Context, transaction *request.TransactionReq) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	limit, err := uc.TransactionRepository.GetLimitForUpdate(tx, transaction.CustomerID, transaction.TenorMonth)
	if err != nil {
		uc.Log.Error("Error fetching customer limit: ", err)
		tx.Rollback()
		return fmt.Errorf("customer don't have limit")
	}

	if limit.LimitAmount-limit.UsedAmount < transaction.OTR {
		tx.Rollback()
		return fmt.Errorf("insufficient limit")
	}

	err = uc.TransactionRepository.Create(tx, transaction)
	if err != nil {
		uc.Log.Error("Error creating transaction: ", err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fiber.ErrConflict
		}
		return err
	}

	limit.UsedAmount += transaction.OTR
	err = uc.TransactionRepository.UpdateLimit(tx, limit)
	if err != nil {
		uc.Log.Error("Error updating customer limit: ", err)
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (uc *TransactionUseCase) GetAllTransactions(ctx context.Context, customerID int64) ([]response.TransactionRes, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	var transactions []entity.Transaction
	err := uc.TransactionRepository.GetAllTransactions(tx, &transactions, customerID)
	if err != nil {
		uc.Log.Error(err)
		return nil, err
	}
	tx.Commit()

	return response.ToTransactionResList(&transactions), nil
}

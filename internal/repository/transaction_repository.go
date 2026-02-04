package repository

import (
	"xyz-multifinance/internal/entity"
	"xyz-multifinance/internal/model/request"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
	Log *logrus.Logger
}

func NewTransactionRepository(log *logrus.Logger) *TransactionRepository {
	return &TransactionRepository{
		Log: log,
	}
}

func (r *TransactionRepository) Create(
	tx *gorm.DB,
	transaction *request.TransactionReq,
) error {
	err := tx.Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *TransactionRepository) GetAllTransactions(
	tx *gorm.DB,
	transactions *[]entity.Transaction,
	customerId int64,
) error {
	return tx.Where("customer_id = ?", customerId).Find(transactions).Error
}

func (r *TransactionRepository) GetLimitForUpdate(
	tx *gorm.DB,
	customerID int64,
	tenor int,
) (*entity.CustomerLimit, error) {

	var limit entity.CustomerLimit

	err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("customer_id = ? AND tenor_month = ?", customerID, tenor).
		First(&limit).Error

	if err != nil {
		return nil, err
	}

	return &limit, nil
}

func (r *TransactionRepository) UpdateLimit(
	tx *gorm.DB,
	limit *entity.CustomerLimit,
) error {
	err := tx.Model(&entity.CustomerLimit{}).
		Where("customer_id = ? AND tenor_month = ?", limit.CustomerID, limit.TenorMonth).
		Update("used_amount", limit.UsedAmount).Error
	if err != nil {
		return err
	}
	return nil
}

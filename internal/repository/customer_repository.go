package repository

import (
	"xyz-multifinance/internal/entity"
	"xyz-multifinance/internal/model/request"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Repository[entity.Customer]
	Log *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomerRepository {
	return &CustomerRepository{
		Log: log,
	}
}

func (r *CustomerRepository) Create(
	tx *gorm.DB,
	customer *request.CustomerReq,
) error {
	return tx.Create(&customer).Error
}

func (r *CustomerRepository) Update(
	tx *gorm.DB,
	customer *entity.Customer,
) error {
	return tx.Save(customer).Error
}

func (r *CustomerRepository) FindByNik(
	tx *gorm.DB,
	user *entity.Customer,
	id string,
) error {
	return tx.Where("nik = ?", id).First(user).Error
}

func (r *CustomerRepository) FindByID(
	tx *gorm.DB,
	user *entity.Customer,
	id int64,
) error {
	return tx.Where("id = ?", id).Take(user).Error
}

func (r *CustomerRepository) FindByToken(
	tx *gorm.DB,
	user *entity.Customer,
	token string,
) error {
	r.Log.Info("Finding user by token: ", token)
	return tx.Where("token = ?", token).Take(user).Error
}

package test

import (
	"testing"

	"xyz-multifinance/internal/entity"

	"github.com/stretchr/testify/assert"
)

func GetFirstCustomer(t *testing.T) *entity.Customer {
	customer := new(entity.Customer)
	err := db.First(customer).Error
	assert.Nil(t, err)
	return customer
}

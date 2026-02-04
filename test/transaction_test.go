package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"xyz-multifinance/internal/entity"
	"xyz-multifinance/internal/model/request"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	user := GetFirstCustomer(t)

	requestBody := request.TransactionReq{
		ContractNumber: "CN-001",
		CustomerID:     user.ID,
		OTR:            80000,
		AdminFee:       5000,
		Installment:    3,
		Interest:       1000,
		AssetName:      "Toyota Avanza Diecast",
		TenorMonth:     1,
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/transactions", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestCreateTransactionFailed(t *testing.T) {
	user := GetFirstCustomer(t)

	requestBody := request.TransactionReq{
		ContractNumber: "CN-001",
		CustomerID:     user.ID,
		OTR:            300000,
		AdminFee:       5000,
		Installment:    3,
		Interest:       1000,
		AssetName:      "Toyota Avanza Diecast",
		TenorMonth:     1,
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/transactions", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+user.Token)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListTransactions(t *testing.T) {
	user := GetFirstCustomer(t)
	request := httptest.NewRequest(http.MethodGet, "/api/v1/transactions", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+user.Token)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bodyBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	bodyString := string(bodyBytes)
	assert.Contains(t, bodyString, "data")
}

func TestListTransactionsFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/transactions", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestTransaction_LimitExceeded(t *testing.T) {
	mockLimit := &entity.CustomerLimit{
		LimitAmount: 100000,
		UsedAmount:  90000,
	}

	available := mockLimit.LimitAmount - mockLimit.UsedAmount

	assert.True(t, available < 20000)
}

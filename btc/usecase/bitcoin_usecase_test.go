package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/essentiaone/ess-atomic-swap/btc/usecase"
	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/stretchr/testify/assert"
)

type ConfirmedTransaction struct{}

func (c *ConfirmedTransaction) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return &models.BitcoinTransaction{
		Confirmations: 6,
	}, nil
}

type NotConfirmedTransaction struct{}

func (c *NotConfirmedTransaction) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return &models.BitcoinTransaction{
		Confirmations: 4,
	}, nil
}

type InvalidRequest struct{}

func (c *InvalidRequest) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return nil, errors.New("invalid response from bitcoin node")
}

type ErrorResponse struct{}

func (c *ErrorResponse) ExecuteRequest(method string, target interface{}, params ...interface{}) (interface{}, error) {
	return "some data", errors.New("invalid response from bitcoin node")
}

func TestCheckTxStatus(t *testing.T) {
	fmt.Println("Transaction confirmed success")
	bitcoin := usecase.New(nil, &ConfirmedTransaction{})
	isSuccess, err := bitcoin.CheckTxStatus("Success")
	assert.NoError(t, err)
	assert.True(t, isSuccess)

	fmt.Println("Transaction not confirmed")
	bitcoin = usecase.New(nil, &NotConfirmedTransaction{})
	isSuccess, err = bitcoin.CheckTxStatus("Failed")
	assert.NoError(t, err)
	assert.False(t, isSuccess)

	fmt.Println("Invalid response from ethereum node")
	bitcoin = usecase.New(nil, &InvalidRequest{})
	isSuccess, err = bitcoin.CheckTxStatus("Invalid")
	assert.Error(t, err)
	assert.False(t, isSuccess)

	fmt.Println("Error response from bictoin node")
	bitcoin = usecase.New(nil, &ErrorResponse{})
	isSuccess, err = bitcoin.CheckTxStatus("Error")
	assert.Error(t, err)
	assert.False(t, isSuccess)
}

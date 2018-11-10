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
	return nil, errors.New("invalid response from Bitcoin node")
}

type ErrorResponse struct{}

func (c *ErrorResponse) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return "some data", errors.New("error response from Bitcoin node")
}

func TestTxStatus(t *testing.T) {
	fmt.Println("Transaction confirmed success")
	bitcoin := usecase.New(&ConfirmedTransaction{})
	isSuccess, err := bitcoin.TxStatus("Success")
	assert.NoError(t, err)
	assert.True(t, isSuccess)

	fmt.Println("Transaction not confirmed")
	bitcoin = usecase.New(&NotConfirmedTransaction{})
	isSuccess, err = bitcoin.TxStatus("Failed")
	assert.NoError(t, err)
	assert.False(t, isSuccess)

	fmt.Println("Invalid response from Bitcoin node")
	bitcoin = usecase.New(&InvalidRequest{})
	isSuccess, err = bitcoin.TxStatus("Invalid")
	assert.Error(t, err)
	assert.False(t, isSuccess)

	fmt.Println("Error response from Bitcoin node")
	bitcoin = usecase.New(&ErrorResponse{})
	isSuccess, err = bitcoin.TxStatus("Error")
	assert.Error(t, err)
	assert.False(t, isSuccess)
}

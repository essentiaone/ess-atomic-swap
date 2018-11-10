package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/essentiaone/ess-atomic-swap/eth/usecase"
	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/stretchr/testify/assert"
)

type ConfirmedTransaction struct{}

func (c *ConfirmedTransaction) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return &models.EthereumTransaction{
		Status: "0x1",
	}, nil
}

type NotConfirmedTransaction struct{}

func (c *NotConfirmedTransaction) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return &models.EthereumTransaction{
		Status: "0x0",
	}, nil
}

type InvalidRequest struct{}

func (c *InvalidRequest) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return nil, errors.New("invalid response from Ethereum node")
}

type ErrorResponse struct{}

func (c *ErrorResponse) ExecuteRequest(
	method string,
	target interface{},
	params ...interface{},
) (interface{}, error) {
	return "some data", errors.New("error response from Ethereum node")
}

func TestTxStatus(t *testing.T) {
	fmt.Println("Transaction confirmed success")
	ethereum := usecase.New(&ConfirmedTransaction{})
	isSuccess, err := ethereum.TxStatus("Success")
	assert.NoError(t, err)
	assert.True(t, isSuccess)

	fmt.Println("Transaction not confirmed")
	ethereum = usecase.New(&NotConfirmedTransaction{})
	isSuccess, err = ethereum.TxStatus("Failed")
	assert.NoError(t, err)
	assert.False(t, isSuccess)

	fmt.Println("Invalid response from Ethereum node")
	ethereum = usecase.New(&InvalidRequest{})
	isSuccess, err = ethereum.TxStatus("Invalid")
	assert.Error(t, err)
	assert.False(t, isSuccess)

	fmt.Println("Error response from Ethereum node")
	ethereum = usecase.New(&ErrorResponse{})
	isSuccess, err = ethereum.TxStatus("Error")
	assert.Error(t, err)
	assert.False(t, isSuccess)
}

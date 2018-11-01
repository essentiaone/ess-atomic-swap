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

func TestCheckTxStatus(t *testing.T) {
	fmt.Println("Transaction confirmed success")
	bitcoin := usecase.New(nil, &ConfirmedTransaction{})
	isSuccess := bitcoin.CheckTxStatus("Success")
	assert.True(t, isSuccess)

	fmt.Println("Transaction not confirmed")
	bitcoin = usecase.New(nil, &NotConfirmedTransaction{})
	isSuccess = bitcoin.CheckTxStatus("Failed")
	assert.False(t, isSuccess)

	fmt.Println("Invalid response from ethereum node")
	bitcoin = usecase.New(nil, &InvalidRequest{})
	isSuccess = bitcoin.CheckTxStatus("Invalid")
	assert.False(t, isSuccess)
}

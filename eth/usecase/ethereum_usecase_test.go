package usecase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/stretchr/testify/assert"
)

var (
	tsSuccess = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"result": { "status": "0x1" }}`)
	}))

	tsFailed = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"result": { "status": "0x0" }}`)
	}))

	tsInvalidResponse = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"result": { status: 0x0 }}`)
	}))

	tsInvalidRequest = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
)

func TestGetTransactionInfo(t *testing.T) {
	fmt.Println("Success response from EHT node. With correct tx info.")
	responseBody, err := getTransactionInfo(tsSuccess.URL, "txHash")
	assert.Empty(t, err)
	response, ok := responseBody.Result.(*models.EthereumTransaction)
	assert.True(t, ok)
	assert.Equal(t, "0x1", response.Status)

	fmt.Println("Success response from EHT node. With invalid tx info.")
	responseBody, err = getTransactionInfo(tsInvalidResponse.URL, "txHash")
	assert.Empty(t, responseBody)
	assert.Error(t, err)

	fmt.Println("Failed response from ETH node.")
	responseBody, err = getTransactionInfo(tsInvalidRequest.URL, "txHash")
	assert.Empty(t, responseBody)
	assert.Error(t, err)
}

func TestCheckTxStatus(t *testing.T) {
	fmt.Println("Transaction fulfilled success")
	eth := New(nil, tsSuccess.URL)
	isSuccess := eth.CheckTxStatus("txHash")
	assert.True(t, isSuccess)

	fmt.Println("Transaction fulfilled failed")
	eth = New(nil, tsFailed.URL)
	isSuccess = eth.CheckTxStatus("txHash")
	assert.False(t, isSuccess)

	fmt.Println("Invalid response from node")
	eth = New(nil, tsInvalidResponse.URL)
	isSuccess = eth.CheckTxStatus("txHash")
	assert.False(t, isSuccess)

	fmt.Println("Invalid request to node")
	eth = New(nil, tsInvalidRequest.URL)
	isSuccess = eth.CheckTxStatus("txHash")
	assert.False(t, isSuccess)
}

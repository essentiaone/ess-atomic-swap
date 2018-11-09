package rpc_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/essentiaone/ess-atomic-swap/models"
	"github.com/essentiaone/ess-atomic-swap/rpc"
	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals
var (
	validateGenerateRequest = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rpcBody := &rpc.RequestRPC{}
		err := json.NewDecoder(r.Body).Decode(rpcBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{"result":{"status":"0x1"}}`)
	}))

	successResponse = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{"result":{"status":"0x1"}}`)
	}))

	invalidResponseFromServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"result":{"status":"invalid_JOSN}}`) // after 'invalid_JSON' absent "
	}))

	emptyResponse = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"result":null}`)
	}))

	errorResponse = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":{"message":"cannot find tx transaction"}}`)
	}))
)

func TestExecuteRequest(t *testing.T) {
	fmt.Println("Correctly generated request to RPC server")
	rpcClient := rpc.New(validateGenerateRequest.URL)
	targetMap := map[string]interface{}{}
	_, err := rpcClient.ExecuteRequest("text_method", targetMap, nil)
	assert.NoError(t, err)

	fmt.Println("Success response from RPC server")
	rpcClient = rpc.New(successResponse.URL)
	targetETH := &models.EthereumTransaction{}
	response, err := rpcClient.ExecuteRequest("test_method", targetETH, nil)
	assert.NoError(t, err)
	transactionInfo, ok := response.(*models.EthereumTransaction)
	assert.True(t, ok)
	assert.Equal(t, "0x1", transactionInfo.Status)

	fmt.Println("Generate invalid request")
	rpcClient = rpc.New(":")
	_, err = rpcClient.ExecuteRequest("test_method", nil, nil)
	assert.Error(t, err)

	fmt.Println("Request to invalid server")
	rpcClient = rpc.New("https://invalid_address")
	_, err = rpcClient.ExecuteRequest("test_method", nil, nil)
	assert.Error(t, err)

	fmt.Println("Invalid response from server")
	rpcClient = rpc.New(invalidResponseFromServer.URL)
	targetMap = map[string]interface{}{}
	_, err = rpcClient.ExecuteRequest("test_method", targetMap, nil)
	assert.Error(t, err)

	fmt.Println("Empty response from server")
	rpcClient = rpc.New(emptyResponse.URL)
	targetMap = map[string]interface{}{}
	_, err = rpcClient.ExecuteRequest("test_method", targetMap, nil)
	assert.Error(t, err)

	fmt.Println("Error response from server")
	rpcClient = rpc.New(errorResponse.URL)
	targetMap = map[string]interface{}{}
	_, err = rpcClient.ExecuteRequest("test_method", targetMap, nil)
	assert.Error(t, err)
	assert.Equal(t, "cannot find tx transaction", err.Error())
}

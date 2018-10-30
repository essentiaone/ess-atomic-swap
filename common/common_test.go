package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ErrorRequestGenerator struct{}

func (g *ErrorRequestGenerator) GenerateRequest() (*http.Request, error) {
	return nil, errors.New("cannot create request")
}

type InvalidRequestGenerator struct{}

func (g *InvalidRequestGenerator) GenerateRequest() (*http.Request, error) {
	return http.NewRequest("POST", "invalid_address", nil)
}

type ErrorStatusGenerator struct {
	TsServer *httptest.Server
}

func (g *ErrorStatusGenerator) GenerateRequest() (*http.Request, error) {
	return http.NewRequest("POST", g.TsServer.URL, nil)
}

type SuccessRequestGenerator struct {
	TsServer *httptest.Server
}

func (g *SuccessRequestGenerator) GenerateRequest() (*http.Request, error) {
	return http.NewRequest("POST", g.TsServer.URL, nil)
}

func TestHTTPRequestToNode(t *testing.T) {
	fmt.Println("Emit error when generating request")
	errorRequest := &ErrorRequestGenerator{}
	responseBody, err := HTTPRequestToNode(errorRequest)
	assert.Error(t, err)
	assert.Nil(t, responseBody)

	fmt.Println("Send request to broken node")
	invalidRequest := &InvalidRequestGenerator{}
	responseBody, err = HTTPRequestToNode(invalidRequest)
	assert.Error(t, err)
	assert.Nil(t, responseBody)

	fmt.Println("Server response error status code")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	errorStatusRequest := &ErrorStatusGenerator{
		TsServer: httptest.NewServer(handler),
	}
	responseBody, err = HTTPRequestToNode(errorStatusRequest)
	assert.Error(t, err)
	assert.Nil(t, responseBody)
	assert.Equal(t, "error status code", err.Error()) // TODO: better create error type in common
	errorStatusRequest.TsServer.Close()

	fmt.Println("Success request to node")
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{"result": "success"}`)
	})
	successRequest := &SuccessRequestGenerator{
		TsServer: httptest.NewServer(handler),
	}
	responseBody, err = HTTPRequestToNode(successRequest)
	assert.NoError(t, err)
	response := map[string]string{}
	err = json.NewDecoder(responseBody).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response, "result")
	assert.Equal(t, "success", response["result"])
}

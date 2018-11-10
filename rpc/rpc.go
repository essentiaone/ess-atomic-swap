package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	nodeError           = "invalid request to node"
	timeWaitResponseRPC = 10 // TODO: in future move to config
)

// RequestRPC structure for body request to RPC server
type RequestRPC struct {
	Params  interface{} `json:"params"`
	Address string      `json:"-"`
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      uint8       `json:"id"`
}

// ResponseErrorRPC structure for error response RPC server
type ResponseErrorRPC struct {
	Message string `json:"message"`
}

// ResponseRPC structure for response RPC server
type ResponseRPC struct {
	Result interface{}       `json:"result,omitempty"`
	Error  *ResponseErrorRPC `json:"error,omitempty"`
}

// New create new RPC client
func New(address string) *RequestRPC {
	return &RequestRPC{
		Address: address,
		ID:      1,
		JSONRPC: "2.0",
	}
}

// ExecuteRequest to RPC server and return 'target' structure
func (r *RequestRPC) ExecuteRequest(method string, target interface{}, params ...interface{}) (interface{}, error) {

	r.Method = method
	r.Params = params

	rawData := []byte{}
	buf := bytes.NewBuffer(rawData)
	err := json.NewEncoder(buf).Encode(r)
	if err != nil {
		return nil, errors.Wrap(err, "cannot encode RPC body")
	}

	request, err := http.NewRequest("POST", r.Address, buf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create RPC request")
	}

	request.Header.Add("Content-type", "application/json")

	client := http.Client{Timeout: time.Second * timeWaitResponseRPC}
	rawResponse, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "cannot send request to RPC server")
	}

	responseRPC := &ResponseRPC{Result: target}
	err = json.NewDecoder(rawResponse.Body).Decode(responseRPC)
	defer func() {
		if rawResponse.Body != nil {
			rawResponse.Body.Close()
		}
	}()
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode response from RPC server")
	}

	/**
	If 'Result' and 'Error' is empty,
	you could not make a request to RPC server.
	If 'Error' not empty,
	RPC server accepted invalid request data
	*/

	if responseRPC.Result == nil && responseRPC.Error == nil {
		return nodeError, errors.New(nodeError)
	} else if responseRPC.Error != nil {
		return nodeError, errors.New(responseRPC.Error.Message)
	}

	return responseRPC.Result, nil
}

package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RequestRPC structure for body request to RPC server
type RequestRPC struct {
	Params  interface{} `json:"params"`
	Address string      `json:"-"`
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      uint8       `json:"id"`
}

// ResponseRPC structure for response RPC server
type ResponseRPC struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
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
		return nil, err
	}

	request, err := http.NewRequest("POST", r.Address, buf)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-type", "application/json")

	client := http.Client{}
	rawResponse, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseRPC := &ResponseRPC{Result: target}
	err = json.NewDecoder(rawResponse.Body).Decode(responseRPC)
	defer func() {
		if rawResponse.Body != nil {
			rawResponse.Body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	return responseRPC.Result, nil
}

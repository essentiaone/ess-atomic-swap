package common

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RequestEthereumBody structure for body request to node
// TODO: use easyjson
type RequestEthereumBody struct {
	Params  interface{} `json:"params"`
	Address string      `json:"-"`
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      uint8       `json:"id"`
}

// ResponseEthereumNode strcut for response ethereum node
type ResponseEthereumNode struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

// GenerateRequest generate request to node
func (r *RequestEthereumBody) GenerateRequest() (*http.Request, error) {
	rawData := []byte{}
	buf := bytes.NewBuffer(rawData)
	err := json.NewEncoder(buf).Encode(r)
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", r.Address, bytes.NewBuffer(buf.Bytes()))
}

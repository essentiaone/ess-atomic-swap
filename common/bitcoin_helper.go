package common

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RequestBitcoinBody structure for body request to node
type RequestBitcoinBody struct {
	Params  interface{} `json:"params"`
	Address string      `json:"-"`
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `josn:"method"`
	ID      uint8       `json:"id"`
}

// ResponseBitcoinNode strcut for response Bitcoin node
type ResponseBitcoinNode struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

// GenerateRequest generate request to node
func (r *RequestBitcoinBody) GenerateRequest() (*http.Request, error) {
	rawData := []byte{}
	buf := bytes.NewBuffer(rawData)
	err := json.NewEncoder(buf).Encode(r)
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", r.Address, bytes.NewBuffer(buf.Bytes()))
}

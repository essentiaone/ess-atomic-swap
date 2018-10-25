package common

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: implement codecov for all unit tests
func TestGenerateRequest(t *testing.T) {
	fmt.Println("Generate success request")
	rawRequest := &RequestEthereumBody{
		Params:  []interface{}{"param1", "param2", "param3"},
		Address: "http://google.com",
		JSONRPC: "2.0",
		Method:  "Steal_personal_data",
		ID:      1,
	}
	request, err := rawRequest.GenerateRequest()
	assert.NoError(t, err)
	assert.Equal(t, "POST", request.Method)
	requestBody := &RequestEthereumBody{}
	err = json.NewDecoder(request.Body).Decode(requestBody)
	assert.NoError(t, err)
	rawRequest.Address = "" // because 'Address' not participating in serialize/deserialize JSON
	assert.Equal(t, rawRequest, requestBody)
}

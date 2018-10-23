package common

import (
	"io"
	"net/http"
)

// RequestGenerator generate request to node
type RequestGenerator interface {
	GenerateRequest() (*http.Request, error)
}

// HTTPRequestToNode implement http request to node
// Теститься тоже изи
func HTTPRequestToNode(r RequestGenerator) (io.ReadCloser, error) {
	request, err := r.GenerateRequest()
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= http.StatusBadRequest {
		return nil, err // TODO add wrapper for error
	}
	return response.Body, nil
}

package common

import (
	"errors"
	"io"
	"net/http"
)

// RequestGenerator generate request to node
type RequestGenerator interface {
	GenerateRequest() (*http.Request, error)
}

// HTTPRequestToNode implement http request to node
func HTTPRequestToNode(r RequestGenerator) (io.ReadCloser, error) {
	request, err := r.GenerateRequest()
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("error status code")
	}
	return response.Body, nil
}

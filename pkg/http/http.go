package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPConfig defines the configuration for a request
type HTTPConfig struct {
	URL     string
	Payload interface{}
	Token   string
}

// HTTPResult defines the result of a request
type HTTPResult struct {
	Err    error
	Body   []byte
	Status int
}

// Get returns a GET request
func Get(config *HTTPConfig) (*HTTPResult, error) {
	return HandleRequest("GET", config)
}

// Post returns a POST request
func Post(config *HTTPConfig) (*HTTPResult, error) {
	return HandleRequest("POST", config)
}

// Put returns a PUT request
func Put(config *HTTPConfig) (*HTTPResult, error) {
	return HandleRequest("PUT", config)
}

// Delete returns a DELETE request
func Delete(config *HTTPConfig) (*HTTPResult, error) {
	return HandleRequest("DELETE", config)
}

func HandleRequest(method string, config *HTTPConfig) (*HTTPResult, error) {
	client := &http.Client{}

	// Prepare body
	reqBody := make([]byte, 0)
	if config.Payload != nil {
		bytes, err := json.Marshal(config.Payload)
		if err != nil {
			return nil, err
		}
		reqBody = bytes
	}

	// Create request
	req, err := http.NewRequest(method, config.URL, bytes.NewBuffer(reqBody))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	// Add authorization
	addAuthorization(req, config)

	resp, err := client.Do(req)
	result := &HTTPResult{
		Err:    err,
		Status: resp.StatusCode,
	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result.Body = body
	}

	return result, nil
}

func addAuthorization(req *http.Request, config *HTTPConfig) {
	if config.Token == "" {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))
}

// Successful indicates if the request was successful or not
func (result *HTTPResult) Successful() bool {
	return result.Status < 400
}

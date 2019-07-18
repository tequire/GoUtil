package identity

import (
	"fmt"

	"github.com/tequire/GoUtil/pkg/http"
)

func handleRequest(client *Client, method string, endpoint string, payload interface{}) (*http.HTTPResult, error) {
	result, err := http.HandleRequest(method, &http.HTTPConfig{
		URL:     fmt.Sprintf("%s/%s", getIdentityServerURL(client.isProd), endpoint),
		Token:   client.token,
		Payload: payload,
	})
	if err != nil {
		return nil, err
	}

	if !result.Successful() {
		return nil, fmt.Errorf("Response returned %d - Error: %s", result.Status, string(result.Body))
	}
	return result, nil
}

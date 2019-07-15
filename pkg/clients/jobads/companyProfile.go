package jobads

import (
	"encoding/json"
	"fmt"

	"github.com/tequire/GoUtil/pkg/config"
	"github.com/tequire/GoUtil/pkg/http"
)

// GetAllArticles gets all the articles
func (c *Client) GetAllArticles() ([]*Article, error) {
	result, err := handleGetProfile(c, "api/v1/organization/article?limit=2147483647") // Fetches with a limit of MAX_INT
	if err != nil {
		return nil, err
	}

	var res struct {
		Items []*Article
	}

	err = json.Unmarshal(result.Body, &res)
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

// GetAllVideos gets all the videos
func (c *Client) GetAllVideos() ([]*Video, error) {
	result, err := handleGetProfile(c, "api/v1/organization/video?limit=2147483647") // Fetch with a limit of MAX_INT
	if err != nil {
		return nil, err
	}

	var res struct {
		Items []*Video
	}

	err = json.Unmarshal(result.Body, &res)
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

func handleGetProfile(client *Client, endpoint string) (*http.HTTPResult, error) {
	result, err := http.Get(&http.HTTPConfig{
		URL:   fmt.Sprintf("%s/%s", getProfileAPIURL(client.isProd), endpoint),
		Token: client.token,
	})
	if err != nil {
		return nil, err
	}

	if !result.Successful() {
		return nil, fmt.Errorf("Response returned %d - Error: %s", result.Status, string(result.Body))
	}
	return result, nil
}

func getProfileAPIURL(isProd bool) string {
	if isProd {
		return config.ProdHigheredAPI
	}
	return config.DevHigheredAPI
}

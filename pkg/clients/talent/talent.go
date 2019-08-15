package talent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tequire/GoUtil/pkg/config"
	"github.com/tequire/GoUtil/pkg/http"
)

// Client to access the CTS API
type Client struct {
	token  string
	isProd bool
}

// New creates a new CTS client
func New(token string) *Client {
	return &Client{
		token:  token,
		isProd: false,
	}
}

// GetPostingByJobAdID gets a jobad from CTS by a JobAdID
func (c *Client) GetPostingByJobAdID(jobAdID int) (*Job, error) {
	url := fmt.Sprintf("api/v1/posting/%d?isJobAdId=true", jobAdID)

	result, err := handleGet(c, url)
	if err != nil {
		return nil, err
	}

	job := &Job{}
	err = json.Unmarshal(result.Body, job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

// GetJobViewsByIDS gets JobViews from the TalentAPI based on input JobAdIds
func (c *Client) GetJobViewsByIDS(jobAdIds []string) ([]*JobView, error) {
	var sb strings.Builder
	for _, id := range jobAdIds {
		sb.WriteString(fmt.Sprintf("id=%s&", id))
	}

	result, err := handleGet(c, fmt.Sprintf("api/v1/postings%s", sb.String()))
	if err != nil {
		return nil, err
	}

	var jobViews []*JobView
	err = json.Unmarshal(result.Body, &jobViews)
	return jobViews, err
}

func handleGet(client *Client, endpoint string) (*http.HTTPResult, error) {
	result, err := http.Get(&http.HTTPConfig{
		URL:   fmt.Sprintf("%s/%s", getAPIURL(client.isProd), endpoint),
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

func getAPIURL(isProd bool) string {
	if isProd {
		return config.ProdHigheredAPI
	}
	return config.DevHigheredAPI
}

package jobads

import (
	"encoding/json"
	"fmt"
)

// GetAllJobAds gets all the jobads
func (c *Client) GetAllJobAds() ([]*Job, error) {
	result, err := handleGet(c, "api/JobAds")
	if err != nil {
		return nil, err
	}

	var jobAds []*Job
	err = json.Unmarshal(result.Body, &jobAds)
	if err != nil {
		return nil, err
	}
	return jobAds, nil
}

// GetJobAdByID gets an jobAd by ID
func (c *Client) GetJobAdByID(jobAdID int) (*Job, error) {
	result, err := handleGet(c, fmt.Sprintf("api/JobAds/%d", jobAdID))
	if err != nil {
		return nil, err
	}

	jobAd := &Job{}
	err = json.Unmarshal(result.Body, jobAd)
	if err != nil {
		return nil, err
	}
	return jobAd, nil
}

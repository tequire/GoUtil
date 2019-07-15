package jobads

import (
	"encoding/json"
	"fmt"

	"github.com/tequire/GoUtil/pkg/config"
	"github.com/tequire/GoUtil/pkg/http"
)

// Client to access the organization API
type Client struct {
	token  string
	isProd bool
}

// New creates a new organization client
func New(token string) *Client {
	return &Client{
		token:  token,
		isProd: false,
	}
}

// SetProd sets if the client should request against the prod or dev environment
func (c *Client) SetProd(isProd bool) {
	c.isProd = isProd
}

// GetOrganizations gets the organizations at /api/v2/organizations
func (c *Client) GetOrganizations() ([]*Organization, error) {
	result, err := handleGet(c, "api/v2/organizations")
	if err != nil {
		return nil, err
	}

	var organizations []*Organization
	err = json.Unmarshal(result.Body, &organizations)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

// GetOrganizationSettings gets the settings for all organizations
func (c *Client) GetOrganizationSettings() ([]*OrganizationSetting, error) {
	result, err := handleGet(c, "api/v1/organizationsprofile")
	if err != nil {
		return nil, err
	}

	var settings []*OrganizationSetting
	err = json.Unmarshal(result.Body, &settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

// GetOrganizationByID gets an organization by id
func (c *Client) GetOrganizationByID(ID int) (*Organization, error) {
	result, err := handleGet(c, fmt.Sprintf("api/v2/organizations/%d", ID))
	if err != nil {
		return nil, err
	}

	var organization Organization
	err = json.Unmarshal(result.Body, &organization)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// GetOrgSettingByDomain returns an organization-setting by domain
func (c *Client) GetOrgSettingByDomain(domain string) (*OrganizationSetting, error) {
	result, err := handleGet(c, fmt.Sprintf("api/v1/organizationsprofile/domain/%s", domain))
	if err != nil {
		return nil, err
	}

	var setting OrganizationSetting
	err = json.Unmarshal(result.Body, &setting)
	if err != nil {
		return nil, err
	}

	return &setting, nil
}

// GetSchools gets all the schools - api/v2/organizations/category/1
func (c *Client) GetSchools() ([]*Organization, error) {
	result, err := handleGet(c, "api/v2/organizations/category/1")
	if err != nil {
		return nil, err
	}

	var schools []*Organization
	err = json.Unmarshal(result.Body, &schools)
	if err != nil {
		return nil, err
	}

	return schools, nil
}

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
		return config.ProdJobAdsAPI
	}
	return config.DevJobAdsAPI
}

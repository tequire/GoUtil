package identity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tequire/GoUtil/pkg/config"
)

// GrantType describes a oauth grant type
type GrantType string

// ClientCredentials is the 'client_credentials' grant type
const ClientCredentials GrantType = "client_credentials"

// ResourceOwnerPassword is the 'password' grant type
const ResourceOwnerPassword GrantType = "password"

// RefreshToken is the 'refresh_token' grant type
const RefreshToken GrantType = "refresh_token"

// AuthorizationCode is the 'authorization_code' grant type
const AuthorizationCode GrantType = "authorization_code"

// Token holds the accessToken and the refreshToken
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenConfig describes the config for a token-request
type TokenConfig struct {
	ClientID     string
	ClientSecret string
	Password     string
	Username     string
	RefreshToken string
	Code         string
	Scope        string
	GrantType    GrantType
	IsProd       bool
}

func getIdentityServerURL(isProd bool) string {
	apiURL := config.DevIdentityServer
	if isProd {
		apiURL = config.ProdIdentityServer
	}
	return apiURL
}

func PostEmailsQuery(filters EmailsQueryFilters, accessToken string, env config.EnvironmentEnum) ([]EmailsQueryResult, error) {
	emailsQueryResult := make([]EmailsQueryResult, 0)

	// fmt.Println("Sending post request to http://localhost:5000/user/query")

	// res, err := http.Post(&http.HTTPConfig{
	// 	URL:     "http://localhost:5000/user/query",
	// 	Payload: filters,
	// 	Token:   "",
	// })
	// if err != nil {
	// 	return emailsQueryResult, err
	// }
	// if !res.Successful() {
	// 	return emailsQueryResult, fmt.Errorf("Query received with status code: %d", res.Status)
	// }
	// err = json.Unmarshal(res.Body, &emailsQueryResult)
	filtersBytes, err := json.Marshal(filters)
	if err != nil {
		return emailsQueryResult, err
	}
	// fmt.Println(accessToken)
	// fmt.Println()
	// fmt.Println(fmt.Sprint(identityURL, "/user/query"))
	// fmt.Println()

	req, err := http.NewRequest("POST", fmt.Sprint(env, "/user/query"), bytes.NewReader(filtersBytes))
	if err != nil {
		return emailsQueryResult, err
	}
	req.Header = map[string][]string{
		"Content-Type":   {"application/json"},
		"Content-Length": {strconv.Itoa(len(filtersBytes))},
		"Authorization":  {accessToken},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emailsQueryResult, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return emailsQueryResult, errors.New(strconv.Itoa(resp.StatusCode))
	}
	err = json.Unmarshal(body, &emailsQueryResult)

	return emailsQueryResult, err
}

// GetAccessToken gets an access token from IdentityServer
func GetAccessToken(config *TokenConfig) (*Token, error) {
	identityURL := getIdentityServerURL(config.IsProd)

	var data url.Values
	if config.GrantType == "" || config.GrantType == ClientCredentials {
		data = newClientCredPayload(config)
	} else if config.GrantType == ResourceOwnerPassword {
		data = newPasswordPayload(config)
	} else if config.GrantType == AuthorizationCode {
		data = newAuthCodePayload(config)
	} else if config.GrantType == RefreshToken {
		data = newRefreshPayload(config)
	} else {
		return nil, errors.New("Invalid grant type")
	}

	// Send token request
	encodedData := data.Encode()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/connect/token", identityURL), strings.NewReader(encodedData))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read and parse body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}
	tokens := &Token{}
	err = json.Unmarshal(body, tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func newClientCredPayload(config *TokenConfig) url.Values {
	return url.Values{
		"client_id":     []string{config.ClientID},
		"client_secret": []string{config.ClientSecret},
		"grant_type":    []string{string(ClientCredentials)},
		"scope":         []string{config.Scope},
	}
}

func newPasswordPayload(config *TokenConfig) url.Values {
	return url.Values{
		"client_id":     []string{config.ClientID},
		"client_secret": []string{config.ClientSecret},
		"username":      []string{config.Username},
		"password":      []string{config.Password},
		"grant_type":    []string{string(ResourceOwnerPassword)},
		"scope":         []string{config.Scope},
	}
}

func newAuthCodePayload(config *TokenConfig) url.Values {
	return url.Values{
		"client_id":     []string{config.ClientID},
		"client_secret": []string{config.ClientSecret},
		"code":          []string{config.Code},
		"grant_type":    []string{string(AuthorizationCode)},
		"scope":         []string{config.Scope},
	}
}

func newRefreshPayload(config *TokenConfig) url.Values {
	return url.Values{
		"client_id":     []string{config.ClientID},
		"client_secret": []string{config.ClientSecret},
		"refresh_token": []string{config.RefreshToken},
		"grant_type":    []string{string(RefreshToken)},
		"scope":         []string{config.Scope},
	}
}

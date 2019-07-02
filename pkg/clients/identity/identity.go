package identity

import (
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

// Token holds the accessToken and the refreshToken
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenConfig describes the config for a token-request
type TokenConfig struct {
	ClientID     string
	ClientSecret string
	Scope        string
	GrantType    string
	IsProd       bool
}

func getIdentityServerURL(isProd bool) string {
	apiURL := config.DevIdentityServer
	if isProd {
		apiURL = config.ProdIdentityServer
	}
	return apiURL
}

// GetAccessToken gets an access token from IdentityServer
func GetAccessToken(config *TokenConfig) (*Token, error) {
	identityURL := getIdentityServerURL(config.IsProd)

	if config.GrantType == "" {
		config.GrantType = "client_credentials"
	}

	data := url.Values{
		"client_id":     []string{config.ClientID},
		"client_secret": []string{config.ClientSecret},
		"grant_type":    []string{config.GrantType},
		"scope":         []string{config.Scope},
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

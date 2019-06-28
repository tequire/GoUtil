package identity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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

	body, err := json.Marshal(map[string]string{
		"client_id":     config.ClientID,
		"client_secret": config.ClientSecret,
		"grant_type":    "client_credentials",
		"scope":         config.Scope,
	})
	if err != nil {
		return nil, err
	}

	// Send token request
	resp, err := http.Post(fmt.Sprintf("%s/connect/token", identityURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Read and parse body
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
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

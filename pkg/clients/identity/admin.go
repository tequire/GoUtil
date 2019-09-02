package identity

// This file contains method for calls to Identity endpoints that require either
// Admin-role or identity-scopes. (identity.full_access, identity.read_access...etc)

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tequire/GoUtil/pkg/config"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Client to access the IdentityServer API
type Client struct {
	token  string
	isProd bool
}

// New creates a new identity client
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

// CreateInactiveUser ... /api/identity/recover/create
func (c *Client) CreateInactiveUser(data BaseUserPost) (*ExtendedUser, error) {
	result, err := handleRequest(c, "POST", "api/identity/recover/create", data)
	if err != nil {
		return nil, err
	}

	user := &ExtendedUser{}
	err = json.Unmarshal(result.Body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// VerifySchool ... /api/identity/recover/verify
func (c *Client) VerifySchool(data BaseVerifyPost) (*UserSchool, error) {
	result, err := handleRequest(c, "POST", "api/identity/recover/verify", data)
	if err != nil {
		return nil, err
	}

	userSchool := &UserSchool{}
	err = json.Unmarshal(result.Body, userSchool)
	if err != nil {
		return nil, err
	}
	return userSchool, nil
}

// CustomUserToken ... /api/identity/token
func (c *Client) CustomUserToken(userID string, hoursToLast int) (*BaseToken, error) {
	var reqBody struct {
		UserID      string `json:"userId"`
		HoursToLast int    `json:"hoursToLast"`
	}
	reqBody.UserID = userID
	reqBody.HoursToLast = hoursToLast

	result, err := handleRequest(c, "POST", "api/identity/token", reqBody)
	if err != nil {
		return nil, err
	}

	token := &BaseToken{}
	err = json.Unmarshal(result.Body, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ChangeAuthenticationMethod ... /user/activate
func (c *Client) ChangeAuthenticationMethod(data ActiveUserPost) (*User, error) {
	result, err := handleRequest(c, "POST", "user/activate", data)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(result.Body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func PostEmailsQuery(filters EmailsQueryFilters, accessToken string, env config.EnvironmentEnum) ([]EmailsQueryResult, error) {
	emailsQueryResult := make([]EmailsQueryResult, 0)

	filtersBytes, err := json.Marshal(filters)
	if err != nil {
		return emailsQueryResult, err
	}

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
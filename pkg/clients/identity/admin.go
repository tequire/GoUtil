package identity

import "encoding/json"

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

package identity

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
func (c *Client) CreateInactiveUser() {

}

// VerifySchool ... /api/identity/recover/verify
func (c *Client) VerifySchool() {

}

// CustomUserToken ... /api/identity/token
func (c *Client) CustomUserToken() {
	var _ struct {
		UserID      string `json:"userId"`
		HoursToLast int    `json:"hoursToLast"`
	}
}

// ChangeAuthenticationMethod ... /user/activate
func (c *Client) ChangeAuthenticationMethod() {

}

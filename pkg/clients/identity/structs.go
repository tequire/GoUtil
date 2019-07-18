package identity

// BaseUserPost defines the request for CreateInactiveUser
type BaseUserPost struct {
	StudentEmail string `json:"studentEmail"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	OldUserID    int    `json:"oldUserId"`

	Nationality        string `json:"nationality"`
	Avatar             string `json:"avatar"`
	PrimarySchool      *int   `json:"primarySchool,omitempty"`
	SchoolRegisteredAt *int   `json:"schoolRegisteredAt,omitempty"`
}

// BaseVerifyPost defines the request for VerifySchool
type BaseVerifyPost struct {
	UserID       string `json:"userId"`
	SchoolID     int    `json:"schoolId"`
	StudentEmail string `json:"studentEmail"`
}

// ActiveUserPost defines the request for ChangeAuthenticationMethod
type ActiveUserPost struct {
	Token string `json:"token"`

	Provider    string `json:"provider,omitempty"`
	AuthCode    string `json:"authCode,omitempty"`
	RedirectURL string `json:"redirectUrl,omitempty"`

	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	EmailReturnURL string `json:"emailReturnUrl,omitempty"`
}

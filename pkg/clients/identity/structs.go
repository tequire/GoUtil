package identity

// BaseUserPost defines the request for CreateInactiveUser
type BaseUserPost struct {
	StudentEmail string `json:"studentEmail"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	OldUserID    int    `json:"oldUserId"`
	PhoneNumber  string `json:"phoneNumber"`

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

// EmailsQueryFilters defines the possible filters Identity takes for an export emails query
type EmailsQueryFilters struct {
	NewsLetterConsent *bool
	SchoolIDs         *[]int
	UserIDs           *[]string
	Nationality       *string
	// OrderBy           string
}

// EmailsQueryResult defines output result from an export emails query
type EmailsQueryResult struct {
	ID                string
	Email             string
	FirstName		  string
	LastName 		  string
	PrimarySchool     *int
	SchoolIds         []int
	NewsLetterConsent bool
	Nationality       string
	Info              string
}

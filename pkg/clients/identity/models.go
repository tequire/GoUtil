package identity

import (
	"time"

	"github.com/google/uuid"
)

// User defines a user
type User struct {
	ID          *uuid.UUID `json:"id"`
	Email       string     `json:"email,omitempty"`
	UserName    string     `json:"userName,omitempty"`
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	FirstName   string     `json:"firstName,omitempty"`
	LastName    string     `json:"lastName,omitempty"`
}

// ExtendedUser defines a user with extended data
type ExtendedUser struct {
	User
	Avatar                 string `json:"avatar,omitempty"`
	Nationality            string `json:"nationality,omitempty"`
	NewsLetterConsent      bool   `json:"newsLetterConsent,omitempty"`
	MaxNewsletterFrequency int    `json:"maxNewsletterFrequency,omitempty"`
}

// UserSchool defines the connection between a user and a school
type UserSchool struct {
	ID           *uuid.UUID `json:"id"`
	StudentEmail string     `json:"studentEmail"`
	IsVerified   bool       `json:"isVerified"`
	SchoolID     int        `json:"schoolId"`
	CreatedAt    time.Time  `json:"createdAt"`
	VerifiedAt   time.Time  `json:"verifiedAt"`
	UserID       string     `json:"userId"`
}

// BaseToken defines a token-response for the CustomToken-endpoint
type BaseToken struct {
	Token string `json:"token,omitempty"`
}

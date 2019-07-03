package gin

import (
	"errors"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tequire/GoUtil/pkg/auth"
)

// UserInContext is a const for user reference in context
const UserInContext = "USER"

// TokenInContext is a const for token reference in context
const TokenInContext = "TOKEN"

// GetUser gets user from context
func GetUser(ctx *gin.Context) (*User, error) {
	userInterface, exists := ctx.Get(UserInContext)
	if !exists {
		return nil, errors.New("user missing in request")
	}
	user, ok := userInterface.(*User)
	if !ok {
		return nil, errors.New("invalid user object")
	}
	var id uuid.UUID = *user.ID
	return &User{ID: &id}, nil
}

// GetClaim gets the claim by name that is inside a token stored in the gin.Context
func GetToken(ctx *gin.Context, name string) (*oidc.IDToken, error) {
	// Get token
	tokenInterface, exists := ctx.Get(TokenInContext)
	if !exists {
		return nil, errors.New("token missing in context")
	}
	token, ok := tokenInterface.(*oidc.IDToken)
	if !ok {
		return nil, errors.New("invalid token object")
	}
	return token
}

// IsAdminOrAuthorized checks if the user is authorized to a resource or has admin privilages
func IsAdminOrAuthorized(ctx *gin.Context, ownerUUID uuid.UUID, resourceUUID uuid.UUID) bool {
	// Check if resource is authorized
	authorized := ownerUUID.String() == resourceUUID.String()
	if authorized {
		return true
	}

	// Check if admin
	tokenInterface, exists := ctx.Get(TokenInContext)
	if !exists {
		return false
	}
	token, ok := tokenInterface.(*oidc.IDToken)
	if !ok {
		return false
	}
	return auth.AdminPolicy(token)
}

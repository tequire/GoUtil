package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

var verifier *oidc.IDTokenVerifier

// User represents a user from the token.
type User struct {
	ID *uuid.UUID `json:"id"`
}

func authorize(ctx *gin.Context) (*oidc.IDToken, error) {
	header := ctx.GetHeader("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid authorization header")
	}
	return verifier.Verify(ctx, parts[1])
}

func tokenToUser(token *oidc.IDToken) (*User, error) {
	if token == nil {
		return nil, fmt.Errorf("Token is nil")
	}

	id, err := uuid.Parse(token.Subject)
	if err != nil {
		return nil, err
	}
	user := User{
		ID: &id,
	}
	return &user, nil
}

// IsAuthorized checks wether a user is authorized.
func IsAuthorized(policies ...Policy) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Validate token
		token, err := authorize(ctx)
		if err != nil {
			fmt.Println(err.Error())
			ctx.Abort()
			ctx.JSON(403, gin.H{})
			return
		}

		// Validate policies
		for _, policy := range policies {
			if valid := policy(token); !valid {
				ctx.Abort()
				ctx.JSON(403, gin.H{})
				return
			}
		}

		// Get user from token
		user, err := tokenToUser(token)
		if err != nil {
			fmt.Println(err.Error())
			ctx.Abort()
			ctx.JSON(403, gin.H{})
		}
		ctx.Set(UserInContext, user)
		ctx.Set(TokenInContext, token)
	}
}

// SetVerifier sets the token verifier.
func SetVerifier(v *oidc.IDTokenVerifier) {
	verifier = v
}

func init() {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://identity.gethighered.global")
	if err != nil {
		panic(err.Error())
	}
	verifier = provider.Verifier(&oidc.Config{
		ClientID: "UserAPI",
	})
}

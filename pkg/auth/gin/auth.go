package gin

import (
	"fmt"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/tequire/GoUtil/pkg/auth"
)

// User represents a user from the token.
type User struct {
	ID *uuid.UUID `json:"id"`
}

// IsAuthorized checks wether a user is authorized.
func IsAuthorized(policies ...Policy) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		err := auth.Authorize(ctx, header, auth.Verifer(), policies...)
		if err != nil {
			fmt.Println(err.Error())
			ctx.Abort()
			ctx.JSON(403, gin.H{})
			return
		}

		// Get user from token
		user, err := tokenToUser(token)
		if err != nil {
			fmt.Println(err.Error())
			ctx.Abort()
			ctx.JSON(403, gin.H{})
			return
		}
		ctx.Set(UserInContext, user)
		ctx.Set(TokenInContext, token)
	}
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

package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserInContext is a const for user reference
const UserInContext = "USER"

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

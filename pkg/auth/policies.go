package auth

import (
	"github.com/coreos/go-oidc"
)

// UserAccessPolicy is a token policy that requires the scope 'API_FULL_USER_ACCESS'.
func UserAccessPolicy(token *oidc.IDToken) bool {
	return requireScope(token, "API_FULL_USER_ACCESS")
}

package auth

import (
	"github.com/coreos/go-oidc"
)

// Policy is a method that validates a token based on a custom policy
type Policy func(token *oidc.IDToken) bool

type scope struct {
	Scopes []string `json:"scope"`
}

// RequireScope checks if a given scope is contained in a given token.
func requireScope(token *oidc.IDToken, scopes ...string) bool {
	// Read scopes
	scope := scope{}
	token.Claims(&scope)

	// Put scopes in map
	scopesMap := map[string]bool{}
	for _, scope := range scope.Scopes {
		scopesMap[scope] = true
	}

	// Check required scopes
	for _, scope := range scopes {
		if _, ok := scopesMap[scope]; !ok {
			return false
		}
	}
	return true
}

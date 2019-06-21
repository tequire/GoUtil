package auth

import (
	"github.com/coreos/go-oidc"
)

// Policy is a method that validates a token based on a custom policy
type Policy func(token *oidc.IDToken) bool

type scope struct {
	Scopes []string `json:"scope"`
}

type role struct {
	Roles []string `json:"role"`
}

// RequireScope checks if a given scope is contained in a given token.
func requireScope(token *oidc.IDToken, scopes ...string) bool {
	// Read scopes
	readScopes := scope{}
	token.Claims(&readScopes)

	// Put scopes in map
	scopesMap := map[string]bool{}
	for _, scope := range readScopes.Scopes {
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

func requireRole(token *oidc.IDToken, roles ...string) bool {
	// Read roles
	readRoles := role{}
	token.Claims(&readRoles)

	// Put roles in map
	roleMap := map[string]bool{}
	for _, role := range readRoles.Roles {
		roleMap[role] = true
	}

	// Check required scopes
	for _, role := range roles {
		if _, ok := roleMap[role]; !ok {
			return false
		}
	}
	return true
}

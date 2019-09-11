package auth

import (
	"github.com/coreos/go-oidc"
)

// Policy is a method that validates a token based on a custom policy
type Policy func(token *oidc.IDToken) bool

// operation defines there should be an AND or OR operation on the claims
type operation uint8

var and = operation(0)
var or = operation(1)

type scope struct {
	Scopes []string `json:"scope"`
}

type role struct {
	Roles []string `json:"role"`
}

// RequireScope checks if a given scope is contained in a given token.
func requireScope(token *oidc.IDToken, op operation, scopes ...string) bool {
	// Read scopes
	readScopes := scope{}
	token.Claims(&readScopes)

	// Put scopes in map
	scopesMap := map[string]bool{}
	for _, scope := range readScopes.Scopes {
		scopesMap[scope] = true
	}

	// Count required scopes
	count := 0
	for _, scope := range scopes {
		if scopesMap[scope] {
			count++
		}
	}

	if op == and {
		return count == len(scopes)
	}
	return count >= 1 // Or
}

func requireRole(token *oidc.IDToken, op operation, roles ...string) bool {
	// Read roles
	parsedRoles := role{}
	token.Claims(&parsedRoles)

	// Put roles in map
	roleMap := map[string]bool{}
	if parsedRoles.Roles != nil && len(parsedRoles.Roles) > 0 {
		// If was able to parse array, add roles into map
		for _, role := range parsedRoles.Roles {
			roleMap[role] = true
		}
	} else {
		// If the role-claim was not an array, try to parse as string
		var roleType struct {
			Role string `json:"role"`
		}
		token.Claims(&roleType)
		if roleType.Role != "" {
			roleMap[roleType.Role] = true
		}
	}

	// Count required scopes
	count := 0
	for _, scope := range roles {
		if roleMap[scope] {
			count++
		}
	}

	if op == and {
		return count == len(roles)
	}
	return count >= 1 // Or
}

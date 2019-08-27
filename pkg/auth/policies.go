package auth

import (
	"github.com/coreos/go-oidc"
)

// UserAccessPolicy is a token policy that requires the scope 'API_FULL_USER_ACCESS'.
func UserAccessPolicy(token *oidc.IDToken) bool {
	return requireScope(token, and, "API_FULL_USER_ACCESS")
}

// AdminPolicy is a token policy that requires the role 'Admin'
func AdminPolicy(token *oidc.IDToken) bool {
	return requireRole(token, and, "Admin")
}

// HigheredEmployeePolicy is a token policy that requires the role 'HigherEdEmployee'
func HigheredEmployeePolicy(token *oidc.IDToken) bool {
	return requireRole(token, and, "HigherEdEmployee")
}

// AdminOrEmployeePolicy is a token policy that requires the role 'Admin' or 'HigherEdEmployee'
func AdminOrEmployeePolicy(token *oidc.IDToken) bool {
	return requireRole(token, or, "Admin", "HigherEdEmployee")
}

// API1Policy is a token policy that requires the scopes 'api1' and 'api1.full_access'
func API1Policy(token *oidc.IDToken) bool {
	return requireScope(token, and, "api1", "api1.full_access")
}

// GoTalentPolicy is a token policy that requires the scopes 'gotalent.full_access'
func GoTalentPolicy(token *oidc.IDToken) bool {
	return requireScope(token, and, "gotalent.full_access")
}

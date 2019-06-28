package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/tequire/GoUtil/pkg/config"

	"github.com/tequire/GoUtil/pkg/env"

	"github.com/coreos/go-oidc"
)

// TokenVerifier is the current token verifier
var tokenVerifier *oidc.IDTokenVerifier

// CTX is the authorization context
var CTX context.Context

// VerifierConfig defines a config of creating a oidc.IDTokenVerifier
type VerifierConfig struct {
	Authority string
}

// VerifyAuthToken converts a raw Authorization header to a verified token
func VerifyAuthToken(ctx context.Context, header string, verifier *oidc.IDTokenVerifier) (*oidc.IDToken, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid authorization header")
	}
	return verifier.Verify(ctx, parts[1])
}

// IsAuthorized checks wether a user is authorized.
func IsAuthorized(policies ...Policy) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// Validate token
		authHeader := r.Header.Get("Authorization")
		_, err := Authorized(CTX, authHeader, tokenVerifier, policies...)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

// Authorized valides an authorization header's token and validies it's policies
func Authorized(ctx context.Context, authHeader string, verifier *oidc.IDTokenVerifier, policies ...Policy) (*oidc.IDToken, error) {
	// Validate token
	token, err := VerifyAuthToken(ctx, authHeader, verifier)
	if err != nil {
		return nil, err
	}

	// Validate policies
	for _, policy := range policies {
		if valid := policy(token); !valid {
			return nil, errors.New("unauthorized")
		}
	}

	return token, nil
}

// SetVerifier sets the token verifier.
func SetVerifier(config *VerifierConfig) {
	tokenVerifier = NewVerifier(config)
}

// Verifier gets the current oidc.IDTokenVerifier
func Verifier() *oidc.IDTokenVerifier {
	return tokenVerifier
}

// NewDevVerifier returns a token-verifier for dev
func NewDevVerifier() *oidc.IDTokenVerifier {
	return NewVerifier(&VerifierConfig{
		Authority: config.DevIdentityServer,
	})
}

// NewProdVerifier returns a token-verifier for prod
func NewProdVerifier() *oidc.IDTokenVerifier {
	return NewVerifier(&VerifierConfig{
		Authority: config.ProdIdentityServer,
	})
}

// NewVerifier creates a new token-verifier based on a config
func NewVerifier(config *VerifierConfig) *oidc.IDTokenVerifier {
	provider, err := oidc.NewProvider(CTX, config.Authority)
	if err != nil {
		panic(err.Error())
	}
	return provider.Verifier(&oidc.Config{})
}

func init() {
	CTX = context.Background() // Sets the global context

	// Decide which verifier-config used based on environment
	verifier := NewDevVerifier()
	if env.IsProduction() {
		verifier = NewProdVerifier()
	}
	tokenVerifier = verifier // Set the global token-verifier
}

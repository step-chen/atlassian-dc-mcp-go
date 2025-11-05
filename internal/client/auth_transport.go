package client

import (
	"net/http"
)

// contextKey defines a type for context keys to avoid collisions.
type ContextKey string

// Context keys for passing tokens.
var (
	BitbucketTokenKey  = ContextKey("bitbucket_token")
	JiraTokenKey       = ContextKey("jira_token")
	ConfluenceTokenKey = ContextKey("confluence_token")
)

// TokenAuthTransport is an http.RoundTripper that adds an Authorization header
// with a Bearer token read from the request's context.
type TokenAuthTransport struct {
	// TokenKey is the key used to look up the token in the context.
	TokenKey ContextKey
	// Transport is the underlying http.RoundTripper to be called after adding the header.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface.
func (t *TokenAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, ok := req.Context().Value(t.TokenKey).(string)
	if !ok || token == "" {
		// If no token is found, proceed without an Authorization header.
		return t.transport().RoundTrip(req)
	}

	// Clone the request to avoid modifying the original request.
	newReq := req.Clone(req.Context())
	newReq.Header.Set("Authorization", "Bearer "+token)

	return t.transport().RoundTrip(newReq)
}

// transport returns the underlying transport to use.
func (t *TokenAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

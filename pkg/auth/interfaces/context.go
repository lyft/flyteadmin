package interfaces

import (
	"github.com/coreos/go-oidc"
	"github.com/lyft/flyteadmin/pkg/auth/config"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

//go:generate mockery -name=AuthenticationContext -case=underscore

// This interface is a convenience wrapper object that holds all the utilities necessary to run Flyte Admin behind authentication
// It is constructed at the root server layer, and passed around to the various auth handlers and utility functions/objects.
type AuthenticationContext interface {
	OAuth2Config() *oauth2.Config
	Claims() config.Claims
	OidcProvider() *oidc.Provider
	CookieManager() CookieHandler
	Options() config.OAuthOptions
	GetUserInfoUrl() *url.URL
	GetHttpClient() *http.Client
}
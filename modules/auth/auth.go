package auth

import (
	"log/slog"
	"net/http"
)

type Authenticator interface {
	GetAccessToken() (*AuthResponse, *ErrorResponse, error)
}

type client struct {
	c              *http.Client
	consumerKey    string
	consumerSecret string
	logger         *slog.Logger
}

func NewAuthenticator(
	httpClient *http.Client,
	logger *slog.Logger,
	key, secret string,
) Authenticator {
	return &client{
		c:              httpClient,
		logger:         logger,
		consumerKey:    key,
		consumerSecret: secret,
	}
}

package stkpush

import (
	"log/slog"
	"net/http"
)

type STKPushHandler interface {
	SendRequest(
		stkPushRequest Request,
		bearerToken string,
	) (*STKResponse, error)
}

type client struct {
	client *http.Client
	logger *slog.Logger
}

func NewSTKPushHandler(c *http.Client, l *slog.Logger) STKPushHandler {
	return &client{client: c, logger: l}
}

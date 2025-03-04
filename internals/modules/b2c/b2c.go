package b2c

import (
	"log/slog"
	"net/http"
)

type B2CHandler interface {
	SendB2CRequest(b2cRequest *B2CRequest, bearerToken string)(interface{}, error)
}

type client struct {
	client *http.Client
	logger *slog.Logger
}

func NewB2CHandler(c *http.Client, l *slog.Logger) B2CHandler {
	return &client{
		client: c,
		logger: l,
	}
}
package stkpush

import (
	"log/slog"
	"net/http"
)

type STKPushHandler interface {
	SendSTKPushRequest(
		bearerToken string, businessShortCode string,
		transactionType string, amount string,
		msisdn string, partyB string, accountReference string,
		transactionDesc string,)(*STKResponse, error)
	HandleSTKCallbackResponse(*StkCallBackResponse)
}

type client struct {
	client *http.Client
	logger *slog.Logger
}

func NewSTKPushHandler(c *http.Client, l *slog.Logger) STKPushHandler {
	return &client{client: c, logger: l}
}
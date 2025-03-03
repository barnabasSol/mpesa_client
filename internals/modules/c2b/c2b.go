package c2b

import (
	"log/slog"
	"net/http"
)

type C2BHandler interface {
	RegisterURL(registerDto RegisterURLDto, apiKey string) (*RegisterResponse, error)

	//  endpoint doesnt exist but still here for testing purposes
	ProcessPayment(paymentReq PaymentRequest, accessToken string) (*PaymentResponse, error)
}

type client struct {
	client *http.Client
	logger *slog.Logger
}

func NewC2BHandler(c *http.Client, logger *slog.Logger) C2BHandler {
	return &client{
		client: c,
		logger: logger,
	}
}

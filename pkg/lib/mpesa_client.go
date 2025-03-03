package mpesa

import (
	"net/http"

	"github.com/barnabasSol/mpesa_client/internals/modules/auth"
	"github.com/barnabasSol/mpesa_client/internals/modules/c2b"
	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

type mpesaClient struct {
	BaseURL string
	Auth    auth.Authenticator
	C2B     c2b.C2BHandler
}

func New(
	baseUrl, consumerKey, consumerSecret string,
	c *http.Client,
) *mpesaClient {
	client := new(http.Client)
	if c != nil {
		client = c
	} else {
		client = defaultClient()
	}
	logger := loggerConfig()

	shared.BaseURL = baseUrl

	authClient := auth.NewAuthenticator(
		client,
		logger,
		consumerKey,
		consumerSecret,
	)

	c2bClient := c2b.NewC2BHandler(client, logger)

	return &mpesaClient{
		baseUrl,
		authClient,
		c2bClient,
	}
}

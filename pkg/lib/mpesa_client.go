package mpesa

import (
	"net/http"

	"github.com/barnabasSol/mpesa_client/internals/modules/auth"
	"github.com/barnabasSol/mpesa_client/internals/modules/b2c"
	"github.com/barnabasSol/mpesa_client/internals/modules/c2b"
	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
	"github.com/barnabasSol/mpesa_client/internals/modules/stkpush"
)

type mpesaClient struct {
	Env     Env
	Auth    auth.Authenticator
	C2B     c2b.C2BHandler
	STKPush stkpush.STKPushHandler
	B2C     b2c.B2CHandler
}

func New(
	env Env,
	c *http.Client,
	consumerKey, consumerSecret string,
) *mpesaClient {

	handleEnv(env)

	client := new(http.Client)
	if c != nil {
		client = c
	} else {
		client = defaultClient()
	}

	logger := loggerConfig()

	authClient := auth.NewAuthenticator(
		client,
		logger,
		consumerKey,
		consumerSecret,
	)

	c2bClient := c2b.NewC2BHandler(client, logger)
	stkpushClient := stkpush.NewSTKPushHandler(client, logger)
	b2cClient := b2c.NewB2CHandler(client, logger)
	return &mpesaClient{
		env,
		authClient,
		c2bClient,
		stkpushClient,
		b2cClient,
	}
}

type Env string

const (
	Sandbox Env = "sandbox"
	Prod    Env = "production"
)

func handleEnv(env Env) {
	switch env {
	case Sandbox:
		shared.BaseURL = "https://apisandbox.safaricom.et"
	case Prod:
		shared.BaseURL = "https://api.safaricom.co.ke"
	default:
		panic("invalid environment")
	}
}

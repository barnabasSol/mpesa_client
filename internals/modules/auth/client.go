package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

type client struct {
	c      *http.Client
	logger *slog.Logger
}

type ClientHandler interface {
	GetAccessToken(consumerKey, consumerSecret string) (*AuthResponse, *shared.ErrorResponse, error)
}

func NewClient(httpClient *http.Client, logger *slog.Logger) ClientHandler {
	return &client{
		c:      httpClient,
		logger: logger,
	}
}

func (c *client) GetAccessToken(consumerKey, consumerSecret string) (
	*AuthResponse,
	*shared.ErrorResponse,
	error,
) {
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	route := fmt.Sprintf(
		"%s/v1/token/generate?%s",
		shared.BaseURL,
		params.Encode(),
	)

	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err) // Return error
	}
	req.SetBasicAuth(consumerKey, consumerSecret)
	resp, err := c.c.Do(req)
	if err != nil {
		//c.logger.Error(err.Error())
		return nil, nil, fmt.Errorf("failed reading response: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		//c.logger.Error(err.Error())
		return nil, nil, fmt.Errorf("failed reading body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errResp := new(shared.ErrorResponse)

		err = json.Unmarshal(body, errResp)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}
		return nil, errResp, nil
	}

	authResponse := new(AuthResponse)

	err = json.Unmarshal(body, authResponse)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal auth response: %w", err)
	}

	return authResponse, nil, nil
}

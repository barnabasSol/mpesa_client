package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

func (c *client) GetAccessToken() (
	*AuthResponse,
	*ErrorResponse,
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
	req.SetBasicAuth(c.consumerKey, c.consumerSecret)
	resp, err := c.c.Do(req)
	if err != nil {
		//c.logger.Error(err.Error())
		return nil, nil, fmt.Errorf("failed reading response: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		//c.logger.Error(err.Error())
		return nil, nil, fmt.Errorf("failed reading body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errResp := new(ErrorResponse)

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

package c2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

func (c *client) RegisterURL(
	registerDto RegisterURLDto,
	apiKey string,
) (*RegisterResponse, error) {
	params := url.Values{}
	params.Add("apiKey", apiKey)
	route := fmt.Sprintf(
		"%s/v1/c2b-register-url/register?%s",
		shared.BaseURL,
		params.Encode(),
	)
	dtoJson, err := json.Marshal(registerDto)
	if err != nil {
		return nil, fmt.Errorf("failed marshall to json %v", err)
	}
	req, err := http.NewRequest("POST", route, bytes.NewBuffer(dtoJson))
	if err != nil {
		return nil, fmt.Errorf("failed creating request %v", err)
	}
	registerResponse := new(RegisterResponse)
	resp, err := c.client.Do(req)
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid api key")
	}
	if err != nil {
		err = shared.ReadJSON(resp, registerResponse)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse the response body %v", err)
		}
		return registerResponse, nil
	}
	err = shared.ReadJSON(resp, registerResponse)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse the response body %v", err)
	}
	return registerResponse, nil
}

func (c *client) ProcessPayment(paymentReq PaymentRequest, accessToken string) (*PaymentResponse, error) {
	route := fmt.Sprintf(
		"%s/v1/c2b/payments",
		shared.BaseURL,
	)
	jsonData, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", route, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var paymentResponse PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &paymentResponse, nil
}

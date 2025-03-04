package c2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/barnabasSol/mpesa_client/lib/modules/shared"
)

func (c *client) RegisterURL(
	registerDto RegisterURLDto,
	apiKey string,
) (*RegisterResponse, error) {
	params := url.Values{}
	params.Add("apikey", apiKey)
	route := fmt.Sprintf(
		"%s/v1/c2b-register-url/register?%s",
		shared.BaseURL,
		params.Encode(),
	)
	dtoJson, err := json.Marshal(registerDto)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(dtoJson))
	if err != nil {
		return nil, fmt.Errorf("failed creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	registerResponse := new(RegisterResponse)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid API key")
	}

	err = shared.ReadJSON(resp, registerResponse)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse the response body: %v", err)
	}
	return registerResponse, nil
}

func (c *client) ProcessPayment(
	paymentReq PaymentRequest,
	accessToken string,
) (*PaymentResponse, error) {
	route := fmt.Sprintf(
		"%s/v1/c2b/payments",
		shared.BaseURL,
	)
	jsonData, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	var paymentResponse PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		log.Println(resp)
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &paymentResponse, fmt.Errorf("API error: %s", resp.Status)
	}

	return &paymentResponse, nil
}

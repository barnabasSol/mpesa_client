package stkpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

func (c *client) SendSTKPushRequest(
	stkPushRequest STKPushRequest,
	bearerToken string,
) (*STKResponse, error) {
	route := fmt.Sprintf(
		"%s/mpesa/stkpush/v3/processrequest",
		shared.BaseURL,
	)

	stkPushRequest.MerchantRequestID = "Partner name -{{$guid}}"
	stkPushRequest.Timestamp = time.Now().Format("20060102150405")
	pass := generatePassword(stkPushRequest.BusinessShortCode, stkPushRequest.Timestamp)
	if pass == "" {
		return nil, fmt.Errorf("failed to generate password")
	}
	c.logger.Info(pass)
	stkPushRequest.Password = pass
	stkPushRequest.CallBackURL = shared.CallBackURL

	jsonBody, err := json.Marshal(stkPushRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest("POST", route, bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed reading response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var stkResponse STKResponse
	if err := json.NewDecoder(resp.Body).Decode(&stkResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &stkResponse, nil
}

package stkpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/barnabasSol/mpesa_client/lib/modules/shared"
)

func (c *client) SendRequest(
	stkPushRequest Request,
	bearerToken string,
) (*STKResponse, error) {
	route := fmt.Sprintf(
		"%s/mpesa/stkpush/v3/processrequest",
		shared.BaseURL,
	)

	stkPushRequest.MerchantRequestID = "Partner name -{{$guid}}"
	stkPushRequest.Timestamp = time.Now().Format("20060102150405")
	c.logger.Info(stkPushRequest.Timestamp)
	pass := generatePassword(stkPushRequest.BusinessShortCode, stkPushRequest.Timestamp)
	if pass == "" {
		return nil, fmt.Errorf("failed to generate password")
	}

	stkPushRequest.Password = pass
	stkPushRequest.CallBackURL = shared.CallBackURL

	jsonBody, err := json.Marshal(stkPushRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	log.Println(resp)

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

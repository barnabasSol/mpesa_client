package b2c

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/barnabasSol/mpesa_client/modules/shared"
)

func (c *client) SendB2CRequest(
	b2cRequest *B2CRequest,
	bearerToken string,
) (interface{}, error) {
	route := fmt.Sprintf(
		"%s/mpesa/b2c/v1/paymentrequest",
		shared.BaseURL,
	)

	jsonBody, err := json.Marshal(b2cRequest)
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
		var errorResponse B2CErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, fmt.Errorf("HTTP status %d, but could not decode error JSON: %w", resp.StatusCode, err)
		}
		return &errorResponse, fmt.Errorf("HTTP status %d, error: %s", resp.StatusCode, errorResponse.ErrorMessage)
	}

	var successResponse B2CSuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&successResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding success JSON: %w", err)
	}

	return successResponse, nil
}

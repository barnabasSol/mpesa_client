package stkpush

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

func (c *client) SendSTKPushRequest(
	bearerToken string, businessShortCode string, transactionType string, amount string,
	msisdn string, partyB string, accountReference string, transactionDesc string,
) (*STKResponse, error) {
	route := fmt.Sprintf(
		"%s/stkpush/v3/processrequest",
		shared.BaseURL,
	)

	timestamp := time.Now().Format("20060102150405")
	password := sha256.Sum256([]byte(timestamp + businessShortCode))

	stkpushReq := STKPushRequest{
		MerchantRequestID: "Partner name -{{$guid}}",
		BusinessShortCode: businessShortCode,
		Password:          base64.StdEncoding.EncodeToString(password[:]),
		Timestamp:         timestamp,
		TransactionType:   transactionType,
		Amount:            amount,
		PartyA:            msisdn,
		PartyB:            partyB,
		PhoneNumber:       msisdn,
		TransactionDesc:   transactionDesc,
		CallBackURL:       shared.CallBackURL,
		AccountReference:  accountReference,
	}

	jsonBody, err := json.Marshal(stkpushReq)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var stkResponse STKResponse
	if err := json.NewDecoder(resp.Body).Decode(&stkResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &stkResponse, nil
}

func (c *client) HandleSTKCallbackResponse(*StkCallBackResponse) {

}

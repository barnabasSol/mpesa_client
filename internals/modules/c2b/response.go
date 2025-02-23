package c2b

import "time"

type RegisterResponse struct {
	Header Header `json:"header"`
}

type Header struct {
	ResponseCode    int       `json:"responseCode"`
	ResponseMessage string    `json:"responseMessage"`
	CustomerMessage string    `json:"customerMessage"`
	TimeStamp       time.Time `json:"timeStamp"`
}

type PaymentResponse struct {
	RequestRefID   string `json:"RequestRefID"`
	ResponseCode   string `json:"ResponseCode"`
	ResponseDesc   string `json:"ResponseDesc"`
	TransactionID  string `json:"TransactionID"`
	AdditionalInfo []any  `json:"AdditionalInfo"` // Use appropriate type based on expected structure
}

// "RequestRefID": "29900fe1-ac90-45ce-9443-19eec5f31234",
// "ResponseCode": "0",
// "ResponseDesc": "The service request is processed successfully.",
// "TransactionID": "",
// "AdditionalInfo": []

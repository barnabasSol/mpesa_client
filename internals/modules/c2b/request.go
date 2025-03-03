package c2b

import "time"

type RegisterURLDto struct {
	ShortCode       string `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	CommandID       string `json:"CommandID"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}

type PaymentRequest struct {
	RequestRefID     string          `json:"RequestRefID"`
	CommandID        string          `json:"CommandID"`
	Remark           string          `json:"Remark"`
	ChannelSessionID string          `json:"ChannelSessionID"`
	SourceSystem     string          `json:"SourceSystem"`
	Timestamp        time.Time       `json:"Timestamp"`
	Parameters       []KeyValuePair  `json:"Parameters"`
	ReferenceData    *[]KeyValuePair `json:"ReferenceData,omitempty"`
	Initiator        Initiator       `json:"Initiator"`
	PrimaryParty     Party           `json:"PrimaryParty"`
	ReceiverParty    Party           `json:"ReceiverParty"`
}

type KeyValuePair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Initiator struct {
	IdentifierType     int    `json:"IdentifierType"`
	Identifier         string `json:"Identifier"`
	SecurityCredential string `json:"SecurityCredential"`
	SecretKey          string `json:"SecretKey"`
}

type Party struct {
	IdentifierType int     `json:"IdentifierType"`
	Identifier     string  `json:"Identifier"`
	ShortCode      *string `json:"ShortCode,omitempty"`
}

package stkpush

import (
)

type STKPushRequest struct {
	MerchantRequestID string 		  `json:"MerchantRequestID"`
	BusinessShortCode string 		  `json:"BusinessShortCode"`
	Password 		  string 		  `json:"Password"`
	Timestamp 		  string 		  `json:"Timestamp"`
	TransactionType   string 		  `json:"TransactionType"`
	Amount 			  string 		  `json:"Amount"`
	PartyA 			  string 		  `json:"PartyA"`
	PartyB 			  string 		  `json:"PartyB"`
	PhoneNumber 	  string 		  `json:"PhoneNumber"`
	TransactionDesc   string 		  `json:"TransactionDesc"`
	CallBackURL 	  string 		  `json:"CallBackURL"`
	AccountReference  string 		  `json:"AccountReference"`
	ReferenceData     []ReferenceData `json:"ReferenceData,omitempty"`
}

type ReferenceData struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
package stkpush

type ReferenceData struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Request struct {
	//server generated id, you can leave it empty
	MerchantRequestID string `json:"MerchantRequestID"`
	BusinessShortCode string `json:"BusinessShortCode"`
	//leave it empty for the library to generate the password
	Password string `json:"Password"`
	//library can generate it
	Timestamp       string  `json:"Timestamp"`
	TransactionType string  `json:"TransactionType"`
	Amount          float64 `json:"Amount"`
	PartyA          string  `json:"PartyA"`
	PartyB          string  `json:"PartyB"`
	PhoneNumber     string  `json:"PhoneNumber"`
	//url is part of the library, can be left empty
	CallBackURL      string          `json:"CallBackURL"`
	AccountReference string          `json:"AccountReference"`
	TransactionDesc  string          `json:"TransactionDesc"`
	ReferenceData    []ReferenceData `json:"ReferenceData"`
}

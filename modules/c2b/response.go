package c2b

type RegisterResponse struct {
	Header Header `json:"header"`
}

type Header struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	CustomerMessage string `json:"customerMessage"`
	TimeStamp       string `json:"timeStamp"`
}

type PaymentResponse struct {
	RequestRefID   string `json:"RequestRefID"`
	ResponseCode   string `json:"ResponseCode"`
	ResponseDesc   string `json:"ResponseDesc"`
	TransactionID  string `json:"TransactionID"`
	AdditionalInfo []any  `json:"AdditionalInfo"`
}

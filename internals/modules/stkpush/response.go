package stkpush

type STKResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode 		string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage 	string `json:"CustomerMessage"`
}

type StkCallBackResponse struct {
	Body struct {
		StkCallBack struct {
			MerchantRequestID   string 			 `json:"MerchantRequestID"`
			CheckoutRequestID   string 			 `json:"CheckoutRequestID"`
			ResponseCode 		string 			 `json:"ResponseCode"`
			ResponseDescription string 			 `json:"ResponseDescription"`
			CallBackMetaData 	CallBackMetadata `json:"CallBackMetaData"`
		} `json:"stkCallback"`
	} `json:"Body"`
}


type CallBackMetadata struct {
	Item []struct {
		Name  string      `json:"Name"`
		Value interface{} `json:"Value"`
	} `json:"Item"`
}

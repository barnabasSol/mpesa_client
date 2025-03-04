package b2c

type B2CSuccessResponse struct {
	ConversionID string `json:"ConversionID"`
	OriginatorConversionID string `json:"OriginatorConversionID"`
	ResponseCode string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
}

type B2CErrorResponse struct {
	RequestID string `json:"requestId"`
	ErrorCode string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
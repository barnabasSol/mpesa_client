package shared

type Response[T any] struct {
	IsSuccess *bool   `json:"is_success,omitempty"`
	Message   *string `json:"message"`
	Result    *T      `json:"result"`
}

type ErrorResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
}

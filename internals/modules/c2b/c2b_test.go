package c2b

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/barnabasSol/mpesa_client/internals/modules/shared"
)

func TestRegisterURL(t *testing.T) {
	tests := []struct {
		name           string
		apiKey         string
		registerDto    RegisterURLDto
		statusCode     int
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name:   "successful registration",
			apiKey: "valid_api_key",
			registerDto: RegisterURLDto{
				ShortCode:       "101010",
				ResponseType:    "Completed",
				CommandID:       "RegisterURL",
				ConfirmationURL: "http://mydomain.com/c2b/confirmation",
				ValidationURL:   "http://mydomain.com/c2b/validation",
			},
			statusCode:    http.StatusOK,
			expectedError: false,
		},
		{
			name:   "invalid API key",
			apiKey: "invalid_api_key",
			registerDto: RegisterURLDto{
				ShortCode:       "101010",
				ResponseType:    "Completed",
				CommandID:       "RegisterURL",
				ConfirmationURL: "http://mydomain.com/c2b/confirmation",
				ValidationURL:   "http://mydomain.com/c2b/validation",
			},
			statusCode:     http.StatusUnauthorized,
			expectedError:  true,
			expectedErrMsg: "invalid API key",
		},
		{
			name:   "short code already registered",
			apiKey: "valid_api_key",
			registerDto: RegisterURLDto{
				ShortCode:       "101010",
				ResponseType:    "Completed",
				CommandID:       "RegisterURL",
				ConfirmationURL: "http://mydomain.com/c2b/confirmation",
				ValidationURL:   "http://mydomain.com/c2b/validation",
			},
			statusCode:    http.StatusBadRequest,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("expected POST request, got %s", r.Method)
				}
				if r.URL.Path != "/v1/c2b-register-url/register" {
					t.Errorf("expected path /v1/c2b-register-url/register, got %s", r.URL.Path)
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
				}

				if tt.apiKey == "invalid_api_key" {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				if tt.statusCode == http.StatusBadRequest {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				w.WriteHeader(tt.statusCode)
				if tt.statusCode == http.StatusOK {
					registerResponse := RegisterResponse{
						Header: Header{
							ResponseCode:    200,
							ResponseMessage: "Request processed successfully",
							CustomerMessage: "Request processed successfully",
							TimeStamp:       "2024-02-12T02:20:31.390",
						},
					}
					json.NewEncoder(w).Encode(registerResponse)
				}
			}))
			defer ts.Close()

			shared.BaseURL = ts.URL

			c := &client{
				client: &http.Client{},
			}

			registerResponse, err := c.RegisterURL(tt.registerDto, tt.apiKey)

			if (err != nil) != tt.expectedError {
				t.Errorf("RegisterURL() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Errorf("RegisterURL() error message = %v, expected %v", err, tt.expectedErrMsg)
			}

			if registerResponse != nil && registerResponse.Header.ResponseCode != tt.statusCode {
				t.Errorf("RegisterURL() response code = %d, expected %d", registerResponse.Header.ResponseCode, tt.statusCode)
			}
		})
	}
}

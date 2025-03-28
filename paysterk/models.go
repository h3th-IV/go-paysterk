package paysterk

import (
	"net/http"
	"time"
)

// InitializeTransactionRequest represents the request payload
type InitializeTransactionRequest struct {
	Email    string `json:"email"`
	Amount   int    `json:"amount"` //This should be in lower denomination of your currency.
	Currency string `json:"currency"`
}

// InitializeTransactionResponse represnets the response from Paystack
type InitializeTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

// VerifyTransactionResponse represents the verification response
type VerifyTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Status string `json:"status"`
	} `json:"data"`
}

// PaystackCLient handles API requests
type PaystackCLient struct {
	SecretKey string
	BaseURL   string
	Client    *http.Client
}

// NewClient initializes a PaystackClient
func NewClient(secretKey string) *PaystackCLient {
	return &PaystackCLient{
		SecretKey: secretKey,
		BaseURL:   "https://api.paystack.co",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

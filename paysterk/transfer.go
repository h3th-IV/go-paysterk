package paysterk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TransferIntegration represents the integration details for a transfer.
type TransferIntegration struct {
	ID           int    `json:"id"`
	IsLive       bool   `json:"is_live"`
	BusinessName string `json:"business_name"`
}

// TransferRecipientDetails contains the bank account details of a transfer recipient.
type TransferRecipientDetails struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

// TransferRecipient represents a recipient of a transfer.
type TransferRecipient struct {
	Active        bool                     `json:"active"`
	Currency      string                   `json:"currency"`
	Description   string                   `json:"description"`
	Domain        string                   `json:"domain"`
	Email         string                   `json:"email"`
	ID            int                      `json:"id"`
	Integration   int                      `json:"integration"`
	Metadata      map[string]interface{}   `json:"metadata"`
	Name          string                   `json:"name"`
	RecipientCode string                   `json:"recipient_code"`
	Type          string                   `json:"type"`
	IsDeleted     bool                     `json:"is_deleted"`
	Details       TransferRecipientDetails `json:"details"`
	CreatedAt     string                   `json:"created_at"`
	UpdatedAt     string                   `json:"updated_at"`
}

type TransferRecipientRequest struct {
	Type          string `json:"type"` // 'nuban' for bank accounts
	Name          string `json:"name"`
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	Currency      string `json:"currency"`
	Email         string `json:"email,omitempty"`
}

// TransferRecipientResponse represents the response for a created recipient
type TransferRecipientResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    TransferRecipient `json:"data"`
}

// CreateTransferRecipient creates a transfer recipient
func (c *PaystackCLient) CreateTransferRecipient(req TransferRecipientRequest) (string, error) {
	body, err := c.doRequest("POST", "transferrecipient", req)
	if err != nil {
		return "", err
	}

	var response TransferRecipientResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if !response.Status {
		return "", fmt.Errorf("failed to create recipient: %s", response.Message)
	}

	return response.Data.RecipientCode, nil
}

// TransferSession contains the details of a transfer session.
type TransferSession struct {
	Provider string `json:"provider"`
	ID       string `json:"id"`
}

// TransferData represents the data for a transfer.
type TransferData struct {
	Amount        int                 `json:"amount"`
	Currency      string              `json:"currency"`
	Domain        string              `json:"domain"`
	Failures      interface{}         `json:"failures"`
	ID            int                 `json:"id"`
	Integration   TransferIntegration `json:"integration"`
	Reason        string              `json:"reason"`
	Reference     string              `json:"reference"`
	Source        string              `json:"source"`
	SourceDetails interface{}         `json:"source_details"`
	Status        string              `json:"status"`
	TitanCode     string              `json:"titan_code"`
	TransferCode  string              `json:"transfer_code"`
	TransferredAt string              `json:"transferred_at"`
	Recipient     TransferRecipient   `json:"recipient"`
	Session       TransferSession     `json:"session"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
}

type TransferRequest struct {
	Amount    int    `json:"amount"`
	Recipient string `json:"recipient"`
	Reason    string `json:"reason"`
	Currency  string `json:"currency,omitempty"` // Default: NGN
	Source    string `json:"source"`
}

// TransferResponse represents the response from Paystack after initiating a transfer
type TransferResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    TransferData `json:"data"`
}

// InitiateTransfer sends money to a recipient
func (c *PaystackCLient) InitiateTransfer(req TransferRequest) (*TransferResponse, error) {
	body, err := c.doRequest(http.MethodPost, "transfer", req)
	if err != nil {
		return nil, err
	}

	var response TransferResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// FinalizeTransfer completes a transfer with OTP authentication
func (c *PaystackCLient) FinalizeTransfer(transferCode, otp string) (*TransferResponse, error) {
	payload := map[string]string{
		"transfer_code": transferCode,
		"otp":           otp,
	}
	body, err := c.doRequest(http.MethodPost, "transfer/finalize_transfer", payload)
	if err != nil {
		return nil, err
	}

	var response TransferResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// FetchTransfer retrieves details of a specific transfer
func (c *PaystackCLient) FetchTransfer(id string) (*TransferResponse, error) {
	url := fmt.Sprintf("transfer/%s", id)
	body, err := c.doRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response TransferResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListTransfers retrieves all transfers
func (c *PaystackCLient) ListTransfers() ([]TransferData, error) {
	body, err := c.doRequest(http.MethodGet, "transfer", nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status  bool           `json:"status"`
		Message string         `json:"message"`
		Data    []TransferData `json:"data"`
		Meta    struct {
			Total     int `json:"total"`
			Skipped   int `json:"skipped"`
			PerPage   int `json:"perPage"`
			Page      int `json:"page"`
			PageCount int `json:"pageCount"`
		} `json:"meta"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// Verifytransfers check the Status of a transfer
func (c *PaystackCLient) Verifytransfer(refeence string) (*TransferData, error) {
	url := fmt.Sprintf("transfer/verify/%s", refeence)
	body, err := c.doRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response TransferResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

//Initiate bulk transfer

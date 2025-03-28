package paysterk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InitializeTransactioc starts a new payment process
func (c *PaystackCLient) InitializeTransaction(email string, amount int, currency string) (*InitializeTransactionResponse, error) {
	payload := InitializeTransactionRequest{
		Email:    email,
		Amount:   amount,
		Currency: currency,
	}

	//Make request
	respBody, err := c.doRequest(http.MethodPost, "transaction/initialize", payload)
	if err != nil {
		return nil, err
	}

	//Parse response from API
	var response InitializeTransactionResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Transaction verification
// VerifyTransaction checks the status of a transactioin using the transaction reference
func (c *PaystackCLient) VerifyTransaction(reference string) (*VerifyTransactionResponse, error) {
	//Make request
	endpoint := fmt.Sprintf("transaction/verify/%s", reference)
	body, err := c.doRequest(http.MethodGet, endpoint, reference)
	if err != nil {
		return nil, err
	}

	//Parse response
	var response VerifyTransactionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

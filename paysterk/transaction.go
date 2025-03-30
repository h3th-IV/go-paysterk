package paysterk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TransactionLog represents the log of a transaction.
type TransactionLog struct {
	TimeSpent      int           `json:"time_spent"`
	Attempts       int           `json:"attempts"`
	Authentication string        `json:"authentication"`
	Errors         int           `json:"errors"`
	Success        bool          `json:"success"`
	Mobile         bool          `json:"mobile"`
	Input          []interface{} `json:"input"`
	Channel        string        `json:"channel"`
	History        []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int    `json:"time"`
	} `json:"history"`
}

// TransactionAuthorization represents the authorization details of a transaction.
type TransactionAuthorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	AccountName       string `json:"account_name"`
}

// represent paystack Transaction data
type Transaction struct {
	ID                 int                      `json:"id"`
	Domain             string                   `json:"domain"`
	Status             string                   `json:"status"`
	ReceiptNumber      string                   `json:"receipt_number"`
	Reference          string                   `json:"reference"`
	Amount             int                      `json:"amount"`
	Message            string                   `json:"message"`
	GatewayResponse    string                   `json:"gateway_response"`
	PaidAt             string                   `json:"paid_at"`
	CreatedAt          string                   `json:"created_at"`
	Channel            string                   `json:"channel"`
	Currency           string                   `json:"currency"`
	IPAddress          string                   `json:"ip_address"`
	Metadata           interface{}              `json:"metadata"`
	Log                TransactionLog           `json:"log"`
	Fees               int                      `json:"fees"`
	FeesSplit          string                   `json:"fees_split"`
	Authorization      TransactionAuthorization `json:"authorization"`
	Customer           Customer                 `json:"customer"`
	Plan               map[string]interface{}   `json:"plan"`
	Subaccount         map[string]interface{}   `json:"subaccount"`
	Split              map[string]interface{}   `json:"split"`
	OrderID            string                   `json:"order_id"`
	RequestedAmount    int                      `json:"requested_amount"`
	PosTransactionData map[string]interface{}   `json:"pos_transaction_data"`
}

// Meta represents pagination metadata
type Meta struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	PerPage  int    `json:"perPage"`
}

// TransactionResponse represents the response from Paystack API
type TransactionResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

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
		AccessCode       string `json:"access_code"`
	} `json:"data"`
}

// VerifyTransactionResponse represents the verification response
type VerifyTransactionResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

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
func (c *PaystackCLient) VerifyTransaction(reference string) (*TransactionResponse, error) {
	//Make request
	endpoint := fmt.Sprintf("transaction/verify/%s", reference)
	body, err := c.doRequest(http.MethodGet, endpoint, reference)
	if err != nil {
		return nil, err
	}

	//Parse response
	var response TransactionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type TransactionsResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
	Meta    Meta          `json:"meta"`
}

// FetchAllTransactions retrieves all transactions from Paystack
func (c *PaystackCLient) FetchAllTransactions() (*TransactionsResponse, error) {
	body, err := c.doRequest("GET", "transaction", nil)
	if err != nil {
		return nil, err
	}

	var response TransactionsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// FetchTransactionById retrieves a single
func (c *PaystackCLient) FetchTransactionByID(id int) (*Transaction, error) {
	url := fmt.Sprintf("transaction/%d", id) // ✅ Removed unnecessary curly braces

	body, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response TransactionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

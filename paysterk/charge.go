package paysterk

import (
	"encoding/json"
	"fmt"
)

/*
The Charge API allows merchants to process payments by collecting customer payment details. This is useful when:
- You don't want to redirect customers to Paystack's checkout page.
- You want to charge a customer’s card directly using an authentication flow.
- You want to charge non-card payment methods like bank accounts and mobile money.
*/

// ChargeResponse represents the response from a charge attempt.
type ChargeResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// ChargeData represents the data of a charge response.
type ChargeData struct {
	Amount          int                    `json:"amount"`
	Currency        string                 `json:"currency"`
	TransactionDate string                 `json:"transaction_date"`
	Status          string                 `json:"status"`
	Reference       string                 `json:"reference"`
	Domain          string                 `json:"domain"`
	Metadata        map[string]interface{} `json:"metadata"`
	GatewayResponse string                 `json:"gateway_response"`
	Message         string                 `json:"message"`
	Channel         string                 `json:"channel"`
	IPAddress       string                 `json:"ip_address"`
	Log             string                 `json:"log"`
	Fees            int                    `json:"fees"`
	Authorization   ChargeAuthorization    `json:"authorization"`
	Customer        Customer               `json:"customer"`
	Plan            interface{}            `json:"plan"`
}

// ChargeAuthorization represents the authorization details of a charge.
type ChargeAuthorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

// ChargeRequest represents the request body for initiating a charge
type ChargeRequest struct {
	Email         string                 `json:"email"`
	Amount        int                    `json:"amount"`
	Currency      string                 `json:"currency,omitempty"` // Default: NGN
	Card          *CardDetails           `json:"card,omitempty"`
	Bank          *BankDetails           `json:"bank,omitempty"`
	MobileMoney   *MobileMoneyDetails    `json:"mobile_money,omitempty"`
	Authorization *ChargeAuthorization   `json:"authorization,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// CardDetails represents card payment information
type CardDetails struct {
	Number      string `json:"number"`
	CVV         string `json:"cvv"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
}

// BankDetails represents bank account payment information
type BankDetails struct {
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}

// MobileMoneyDetails represents mobile money payment information
type MobileMoneyDetails struct {
	Phone    string `json:"phone"`
	Provider string `json:"provider"`
}

// Charge initiates a charge for a customer using card, bank, or mobile money
func (c *PaystackCLient) Charge(req ChargeRequest) (*ChargeResponse, error) {
	body, err := c.doRequest("POST", "charge", req)
	if err != nil {
		return nil, err
	}

	var response ChargeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SubmitPIN submits a PIN if required for a charge
func (c *PaystackCLient) SubmitPIN(reference, pin string) (*ChargeResponse, error) {
	payload := map[string]string{
		"reference": reference,
		"pin":       pin,
	}
	body, err := c.doRequest("POST", "charge/submit_pin", payload)
	if err != nil {
		return nil, err
	}

	var response ChargeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SubmitOTP submits an OTP if required for a charge
func (c *PaystackCLient) SubmitOTP(reference, otp string) (*ChargeResponse, error) {
	payload := map[string]string{
		"reference": reference,
		"otp":       otp,
	}
	body, err := c.doRequest("POST", "charge/submit_otp", payload)
	if err != nil {
		return nil, err
	}

	var response ChargeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// VerifyCharge checks the status of a charge using its reference
func (c *PaystackCLient) VerifyCharge(reference string) (*ChargeResponse, error) {
	url := fmt.Sprintf("charge/%s", reference)
	body, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response ChargeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

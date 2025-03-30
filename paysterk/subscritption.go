package paysterk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SubscriptionRequest struct {
	Customer      string `json:"customer"` //Customer's email address or customer code
	PlanCode      string `json:"plan"`
	Authorization string `json:"authorization,omitempty"`
}

// SubscriptionPlan represents a subscription plan.
type SubscriptionPlan struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PlanCode     string `json:"plan_code,omitmepty"`
	Description  string `json:"description,omitempty"`
	Amount       int    `json:"amount"`
	Interval     string `json:"interval"`
	SendInvoices bool   `json:"send_invoices,omitmepty"`
	SendSMS      bool   `json:"send_sms,omitmepty"`
	Currency     string `json:"currency,omitmepty"`
}

// SubscriptionAuthorization represents the authorization details of a subscription.
type SubscriptionAuthorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reuseable         string `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

// SubscriptionCustomer represents the customer information for a subscription.
type SubscriptionCustomer struct {
	FirstName    string                 `json:"first_name"`
	LastName     string                 `json:"last_name"`
	Email        string                 `json:"email"`
	CustomerCode string                 `json:"customer_code"`
	Phone        string                 `json:"phone"`
	Metadata     map[string]interface{} `json:"metadata"`
	RiskAction   string                 `json:"risk_action"`
}

// SubscriptionData represents the data for a subscription.
type SubscriptionData struct {
	ID                int                       `json:"id"`
	Domain            string                    `json:"domain"`
	Status            string                    `json:"status"`
	Quantity          string                    `json:"quantity"`
	SubscriptionCode  string                    `json:"subscription_code"`
	EmailToken        string                    `json:"email_token"`
	Amount            int                       `json:"amount"`
	CronExpression    string                    `json:"cron_expression"`
	NextPaymentDate   string                    `json:"next_payment_date"`
	OpenInvoice       string                    `json:"open_invoice"`
	Plan              SubscriptionPlan          `json:"plan"`
	Authorization     SubscriptionAuthorization `json:"authorization"`
	Customer          SubscriptionCustomer      `json:"customer"`
	Invoices          []interface{}             `json:"invoices"`
	InvoicesHistory   []interface{}             `json:"invoices_history"`
	InvoiceLimit      int                       `json:"invoice_limit"`
	SplitCode         string                    `json:"split_code"`
	MostRecentInvoice interface{}               `json:"most_recent_invoice"`
	CreatedAt         string                    `json:"created_at"`
}

// ExpiringCardSubscription represents the subscription details of an expiring card.
type ExpiringCardSubscription struct {
	ID               int              `json:"id"`
	SubscriptionCode string           `json:"subscription_code"`
	Amount           int              `json:"amount"`
	NextPaymentDate  string           `json:"next_payment_date"`
	Plan             SubscriptionPlan `json:"plan"`
}

// ExpiringCard represents an expiring card.
type ExpiringCard struct {
	ExpiryDate   string                   `json:"expiry_date"`
	Description  string                   `json:"description"`
	Brand        string                   `json:"brand"`
	Subscription ExpiringCardSubscription `json:"subscription"`
	Customer     SubscriptionCustomer     `json:"customer"`
}

// ExpiringCardsEvent represents the event data for expiring cards.
type ExpiringCardsEvent struct {
	Event string         `json:"event"`
	Data  []ExpiringCard `json:"data"`
}

// CreatePlan is used to create a plan on your integration, returns a subscription plan
// Interval is of weekly, monthly, quaterly, biannually and annually.
// Also takes optional
func (c *PaystackCLient) CreatePlan(name, interval, description string, amount int) (*SubscriptionPlan, error) {
	payload := &SubscriptionPlan{
		Name:        name,
		Interval:    interval,
		Amount:      amount,
		Description: description,
	}

	body, err := c.doRequest(http.MethodPost, "plan", payload)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status  bool             `json:"status"`
		Message string           `json:"message"`
		Data    SubscriptionPlan `json:"data"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// FetchPlan retrives details of a plan on your integration
func (c *PaystackCLient) FetchPlan(id, plan_code string) (*SubscriptionPlan, error) {
	endpoint := fmt.Sprintf("plan/%s", plan_code)
	body, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	var response struct {
		Status  bool             `json:"status"`
		Message string           `json:"message"`
		Data    SubscriptionPlan `json:"data"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// UpdatePlan updates a plan details on your integration and returns a message
// Interval is of options: weekly, monthly, quaterly, biannually and annually.
func (c *PaystackCLient) UpdatePlan(plan_code, name, interval string, amount int) (string, error) {
	endpoint := fmt.Sprintf("plan/%s", plan_code)
	planUpdate := &SubscriptionPlan{
		Name:     name,
		Amount:   amount,
		Interval: interval,
	}
	body, err := c.doRequest(http.MethodPut, endpoint, planUpdate)
	if err != nil {
		return "", err
	}
	var response struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Message, nil
}

// CreateSubscription subscribes a customer to a plan, Creates a customer subscription on your integration
func (c *PaystackCLient) CreateSubscription(customerEmail, planCode, authCode string) (*SubscriptionData, error) {
	payload := &SubscriptionRequest{
		Customer:      customerEmail,
		PlanCode:      planCode,
		Authorization: authCode,
	}

	body, err := c.doRequest(http.MethodPost, "subscription", payload)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status  bool             `json:"status"`
		Message string           `json:"message"`
		Data    SubscriptionData `json:"data"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// Get details of a subscription on your integration
func (c *PaystackCLient) FetchSubscription(subscription_code, id string) (*SubscriptionData, error) {
	endpoint := fmt.Sprintf("subscription/%s", subscription_code)

	body, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	var response struct {
		Status  bool             `json:"status"`
		Message string           `json:"message"`
		Data    SubscriptionData `json:"data"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// DisableSubscription stops an active subscription
func (c *PaystackCLient) DisableSubscription(subscriptionCode, emailToken string) error {
	payload := map[string]string{
		"code":  subscriptionCode,
		"token": emailToken,
	}

	_, err := c.doRequest(http.MethodPost, "subscription/disable", payload)
	return err
}

// EnableSubscription reactivates a disabled subscription
func (c *PaystackCLient) EnableSubscription(subscriptionCode, emailToken string) error {
	payload := map[string]string{
		"code":  subscriptionCode,
		"token": emailToken,
	}

	_, err := c.doRequest(http.MethodPost, "subscription/enable", payload)
	return err
}

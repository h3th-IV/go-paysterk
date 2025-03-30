package paysterk

// RefundCustomer represents the customer information for a refund.
type RefundCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// RefundData represents the data for a refund.
type RefundData struct {
	Status               string         `json:"status"`
	TransactionReference string         `json:"transaction_reference"`
	RefundReference      string         `json:"refund_reference"`
	Amount               int            `json:"amount"`
	Currency             string         `json:"currency"`
	Processor            string         `json:"processor"`
	Customer             RefundCustomer `json:"customer"`
	Integration          int            `json:"integration"`
	Domain               string         `json:"domain"`
}

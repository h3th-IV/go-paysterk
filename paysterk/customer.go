package paysterk

// Identification represents the identification details of a customer.
type Identification struct {
	Country       string `json:"country"`
	Type          string `json:"type"`
	BVN           string `json:"bvn,omitempty"`
	Value         string `json:"value,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
}

// CustomerIdentificationData represents the data for a customer identification.
type CustomerIdentificationData struct {
	CustomerID     string         `json:"customer_id"`
	CustomerCode   string         `json:"customer_code"`
	Email          string         `json:"email"`
	Identification Identification `json:"identification"`
	Reason         string         `json:"reason,omitempty"`
}

// Customer represents a customer.
type Customer struct {
	ID                       int                    `json:"id"`
	FirstName                string                 `json:"first_name"`
	LastName                 string                 `json:"last_name"`
	Email                    string                 `json:"email"`
	CustomerCode             string                 `json:"customer_code"`
	Phone                    string                 `json:"phone"`
	Metadata                 map[string]interface{} `json:"metadata"`
	RiskAction               string                 `json:"risk_action"`
	InternationalFormatPhone string                 `json:"international_format_phone"`
}

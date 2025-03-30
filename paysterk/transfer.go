package paysterk

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

package paysterk

// PaymentRequestNotification represents a notification for a payment request.
type PaymentRequestNotification struct {
	SentAt  string `json:"sent_at"`
	Channel string `json:"channel"`
}

// PaymentRequestData represents the data for a payment request.
type PaymentRequestData struct {
	ID               int                          `json:"id"`
	Domain           string                       `json:"domain"`
	Amount           int                          `json:"amount"`
	Currency         string                       `json:"currency"`
	DueDate          string                       `json:"due_date"`
	HasInvoice       bool                         `json:"has_invoice"`
	InvoiceNumber    string                       `json:"invoice_number"`
	Description      string                       `json:"description"`
	PDFURL           string                       `json:"pdf_url"`
	LineItems        []interface{}                `json:"line_items"`
	Tax              []interface{}                `json:"tax"`
	RequestCode      string                       `json:"request_code"`
	Status           string                       `json:"status"`
	Paid             bool                         `json:"paid"`
	PaidAt           string                       `json:"paid_at"`
	Metadata         map[string]interface{}       `json:"metadata"`
	Notifications    []PaymentRequestNotification `json:"notifications"`
	OfflineReference string                       `json:"offline_reference"`
	Customer         int                          `json:"customer"`
	CreatedAt        string                       `json:"created_at"`
}

package paysterk

// DisputeHistory represents the history of a dispute.
type DisputeHistory struct {
	Status    string `json:"status"`
	By        string `json:"by"`
	CreatedAt string `json:"createdAt"`
}

// DisputeMessage represents a message in a dispute.
type DisputeMessage struct {
	Sender    string `json:"sender"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

// DisputeData represents the data for a dispute.
type DisputeData struct {
	ID                   int              `json:"id"`
	RefundAmount         int              `json:"refund_amount"`
	Currency             string           `json:"currency"`
	Status               string           `json:"status"`
	Resolution           string           `json:"resolution"`
	Domain               string           `json:"domain"`
	Transaction          Transaction      `json:"transaction"`
	TransactionReference string           `json:"transaction_reference"`
	Category             string           `json:"category"`
	Customer             Customer         `json:"customer"`
	Bin                  string           `json:"bin"`
	Last4                string           `json:"last4"`
	DueAt                string           `json:"dueAt"`
	ResolvedAt           string           `json:"resolvedAt"`
	Evidence             string           `json:"evidence"`
	Attachments          string           `json:"attachments"`
	Note                 string           `json:"note"`
	History              []DisputeHistory `json:"history"`
	Messages             []DisputeMessage `json:"messages"`
	CreatedAt            string           `json:"created_at"`
	UpdatedAt            string           `json:"updated_at"`
}

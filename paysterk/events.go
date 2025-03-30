package paysterk

// WebhookEvent represents the structure of Paystack webhook payloads
type WebhookEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"` // Generic data field for any event type
}

// CustomerIdentificationEventData represents the data for a customer identification event.
type CustomerIdentificationEventData struct {
	Event string                     `json:"event"`
	Data  CustomerIdentificationData `json:"data"`
}

// DisputeEventData represents the data for a dispute event.
type DisputeEventData struct {
	Event string      `json:"event"`
	Data  DisputeData `json:"data"`
}

// PaymentRequestEventData represents the data for a payment request event.
type PaymentRequestEventData struct {
	Event string             `json:"event"`
	Data  PaymentRequestData `json:"data"`
}

// TransferEventData represents the data for a transfer event.
type TransferEventData struct {
	Event string       `json:"event"`
	Data  TransferData `json:"data"`
}

// TransactionEventData represents the data for a transaction event.
type TransactionEventData struct {
	Event string          `json:"event"`
	Data  Transaction `json:"data"`
}

// RefundEventData represents the data for a refund event.
type RefundEventData struct {
	Event string     `json:"event"`
	Data  RefundData `json:"data"`
}

// SubscriptionEventData represents the data for a subscription event.
type SubscriptionEventData struct {
	Event string           `json:"event"`
	Data  SubscriptionData `json:"data"`
}

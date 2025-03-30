package paysterk

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleWebhook(t *testing.T) {
	// Simulate a webhook payload
	webhookPayload := `{
		"event": "charge.success",
		"data": {
			"id": 123456,
			"reference": "trx_123",
			"amount": 5000,
			"currency": "NGN",
			"customer": {
				"email": "user@example.com"
			},
			"status": "success"
		}
	}`

	req, err := http.NewRequest("POST", "/webhook", bytes.NewBuffer([]byte(webhookPayload)))
	if err != nil {
		t.Fatal(err)
	}

	// Add a fake signature header
	req.Header.Set("x-paystack-signature", "fake_signature")

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WebHookHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, got %v", status)
	}
}

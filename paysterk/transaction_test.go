package paysterk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//NOTE: all keys and references passed here are for test purpose only
// Mock HTTP server
func setupMockServer(response string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	})
	return httptest.NewServer(handler)
}

func TestInitializeTransaction(t *testing.T) {
	mockServer := setupMockServer(`{"status": true, "message": "Success", "data": {"authorization_url": "https://paystack.com/pay/xyz"}}`, 200)
	defer mockServer.Close()

	client := &PaystackCLient{SecretKey: "test_key", BaseURL: mockServer.URL}

	tx, err := client.InitializeTransaction("user@example.com", 5000, "NGN")
	if err != nil || tx.Data.AuthorizationURL == "" {
		t.Errorf("Failed to initialize transaction")
	}
}

func TestVerifyTransaction(t *testing.T) {
	mockServer := setupMockServer(`{"status": true, "message": "Verification successful", "data": {"status": "success"}}`, 200)
	defer mockServer.Close()

	client := &PaystackCLient{SecretKey: "test_key", BaseURL: mockServer.URL}

	resp, err := client.VerifyTransaction("test_ref")
	if err != nil || resp.Data.Status != "success" {
		t.Errorf("Failed to verify transaction")
	}
}
